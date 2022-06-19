package handler

import (
	"context"
	"exl-server/constants"
	"exl-server/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func fileUpload(parentCtx context.Context,  fs service.FileService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancelFunc := context.WithCancel(parentCtx)
		defer cancelFunc()
		vars := mux.Vars(r)
		tenantId := vars["tenant-id"]
		cloudType := vars["cloud-type"]
		log.Printf("TenantId: %s", tenantId)

		contentType := r.Header.Get("Content-Type")
		contentLength := r.Header.Get("Content-Length")
		log.Printf("contentType: %s, contentLength: %s", contentType, contentLength)
		fileExtension, present := constants.ContentTypeToFileExtension[contentType]
		if !present {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Content-Type unsupported"))
			return
		}
		
		cloud := constants.GetCloudType(cloudType)
		if cloud == constants.CloudType("") {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("cloud-type unsupported"))	
			return	
		}
		objectId, err := fs.UploadSync(ctx, tenantId, cloud, r.Body, fileExtension)
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

func MakeFileHandler(ctx context.Context, router *mux.Router, fs service.FileService) {
	router.HandleFunc("/v1/upload", fileUpload(ctx, fs)).Methods("POST").Queries("tenant-id", "{tenant-id:.*}", "cloud-type", "{cloud-type:.*}")
}