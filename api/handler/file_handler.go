package handler

import (
	"context"
	"exl-server/constants"
	"exl-server/service"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func fileUpload(parentCtx context.Context,  fs service.FileService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancelFunc := context.WithCancel(parentCtx)
		defer cancelFunc()
		vars := mux.Vars(r)
		orgId := vars["org-id"]
		cloudType := vars["cloud-type"]
		log.Printf("OrgId: %s", orgId)

		contentType := r.Header.Get("Content-Type")
		contentLength := r.Header.Get("Content-Length")
		log.Printf("contentType: %s, contentLength: %s", contentType, contentLength)
		fileExtension, present := constants.ContentTypeToFileExtension[contentType]
		if !present {
			log.Printf("content-type unsupported\n")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Content-Type unsupported"))
			return
		}
		
		log.Printf("get cloud-type\n")
		cloud := constants.GetCloudType(cloudType)
		if cloud == constants.CloudType("") {
			log.Printf("cloud-type unsupported\n")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("cloud-type unsupported"))	
			return	
		}
		log.Printf("Cloud : %s", cloud)

		objectId, err := fs.UploadSync(ctx, orgId, cloud, r.Body, fileExtension)
		if err != nil {
			if err ==  service.ErrTenantIdNotFound {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(objectId))
	};
}

func fileDownload(parentCtx context.Context,  fs service.FileService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancelFunc := context.WithCancel(parentCtx)
		defer cancelFunc()
		vars := mux.Vars(r)
		objectId := vars["object-id"]
		body, fileMetadata, err := fs.DownloadSync(ctx, objectId)
		if err != nil {
			if err ==  service.ErrObjectIdNotFound {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer body.Close()

		log.Printf("fileName: %s, fileExtension: %s", fileMetadata.Name, fileMetadata.FileExtension)
		contentDisposition := fmt.Sprintf("attachment; filename=%s", fileMetadata.Name)
		contentType, present := constants.FileExtensionToContentType[fileMetadata.FileExtension]
		if !present {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Content Type mismatch"))
			return
		}
		log.Printf("Content-Disposition: %s, Content-Type: %s", contentDisposition, contentType)
		w.Header().Set("Content-Disposition", contentDisposition)
		w.Header().Set("Content-Type", contentType)
		io.Copy(w, body)
	}
}

func MakeFileHandler(ctx context.Context, router *mux.Router, fs service.FileService) {
	router.HandleFunc("/v1/upload", fileUpload(ctx, fs)).Methods("POST").Queries("org-id", "{org-id:.*}", "cloud-type", "{cloud-type:.*}")
	router.HandleFunc("/v1/download", fileDownload(ctx, fs)).Methods("GET").Queries("object-id", "{object-id:.*}")
}