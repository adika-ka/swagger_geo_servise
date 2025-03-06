package handlers

import (
	"net/http"
	"swagger_geo_service/internal/clients"
	"swagger_geo_service/internal/models"
)

func createDaDataClient() *clients.DaDataClient {
	return &clients.DaDataClient{
		ApiKey:     "9c667615626123e3c70123efa6ca12e53ae94e06",
		HttpClient: http.Client{},
	}
}

func handleAPIError(w http.ResponseWriter, err error) {
	if err.Error() == "no addresses found" {
		http.Error(w, "No addresses found", http.StatusNotFound)
	} else {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func convertToPointerSlice(addresses []models.Address) []*models.Address {
	addressPointers := make([]*models.Address, len(addresses))
	for i := range addresses {
		addressPointers[i] = &addresses[i]
	}
	return addressPointers
}
