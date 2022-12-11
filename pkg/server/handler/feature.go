package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/schedule-api/pkg/feature"
)

type FeatureSaver interface {
	Save(ctx context.Context, data feature.Feature) (int, error)
}

func HandleFeatureSave(featureSaver FeatureSaver) http.HandlerFunc {
	type request struct {
		Name string `json:"name"`
	}
	type response struct {
		Id int `json:"id"`
	}

	return func(res http.ResponseWriter, req *http.Request) {
		var body request
		err := json.NewDecoder(req.Body).Decode(&body)
		if err != nil {
			log.Printf("Feature save error - " + err.Error())
			response := &errorResponse{
				Message: err.Error(),
			}
			makeResponse(res, http.StatusBadRequest, response)
			return
		}
		data := feature.Feature{
			Name:        body.Name,
			Responsible: GetUserIdFromRequest(req),
		}
		result, err := featureSaver.Save(req.Context(), data)
		if err != nil {
			log.Printf("Feature save error - " + err.Error())
			response := &errorResponse{
				Message: err.Error(),
			}
			makeResponse(res, http.StatusInternalServerError, response)
			return
		}
		response := &response{
			Id: result,
		}
		log.Printf("Feature save sucess!")
		makeResponse(res, http.StatusOK, response)
	}
}

type FeatureGetter interface {
	GetById(ctx context.Context, id int) (feature.Feature, error)
}

func HandleFeatureGetById(featureGetter FeatureGetter) http.HandlerFunc {
	type response struct {
		Id           int       `json:"id"`
		Name         string    `json:"name"`
		Responsible  int       `json:"responsible"`
		CreationDate time.Time `json:"creationDate"`
		UpdateDate   time.Time `json:"updateDate"`
		Active       bool      `json:"active"`
	}

	return func(res http.ResponseWriter, req *http.Request) {
		idParam, ok := getPathParameterFromRequest(req, "id")
		id, err := strconv.Atoi(idParam)
		if !ok || err != nil {
			log.Printf("Feature get error - id not informed")
			makeResponse(res, http.StatusBadRequest, &errorResponse{Message: "id not informed"})
			return
		}

		result, err := featureGetter.GetById(req.Context(), id)
		if err != nil {
			log.Printf("Feature get error - " + err.Error())
			response := &errorResponse{
				Message: err.Error(),
			}
			makeResponse(res, http.StatusInternalServerError, response)
			return
		}
		response := &response{
			Id:           result.ID,
			Name:         result.Name,
			Responsible:  result.Responsible,
			CreationDate: result.CreationDate,
			UpdateDate:   result.UpdateDate,
			Active:       result.Active,
		}
		log.Printf("Feature get sucess!")
		makeResponse(res, http.StatusOK, response)
	}
}
