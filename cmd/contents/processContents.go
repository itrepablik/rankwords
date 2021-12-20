package contents

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"

	"rankwords/models"
	"rankwords/utils/resp"
	"rankwords/utils/validation"
	"rankwords/utils/validation/schemas"

	logger "github.com/sirupsen/logrus"
)

var errMsg validation.CustomValidationError

// AcceptContentsHandler accepts the contents of the file
func AcceptContentsHandler(w http.ResponseWriter, r *http.Request) {
	// Empty errMsg.ErrorMessage array
	errMsg.ErrorMessage = []string{}
	errInfo := ""

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errInfo = fmt.Sprintf("error reading request body: %s", err.Error())

		logger.WithFields(logger.Fields{
			"error": err,
			"body":  body,
		}).Error(errInfo)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Validate JSON payload
	valErrors, err := validation.SchemaValidation(string(body), schemas.Contents)
	if err != nil {
		errInfo = fmt.Sprintf("error validating request body: %s", err.Error())
		logger.WithFields(logger.Fields{
			"error":     err,
			"valErrors": valErrors,
		}).Error(errInfo)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Parse JSON payload into struct
	contents, err := parseBodyContents(body)
	if err != nil {
		errInfo = fmt.Sprintf("error parsing request body: %s", err.Error())
		logger.WithFields(logger.Fields{
			"error": err,
		}).Error(errInfo)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the top 10 most words from the contents
	topWords := countWords(contents.Text)
	top10Words := rankWords(topWords)

	resp.HttpResponse(w, http.StatusOK, top10Words)
}

// rankWords ranks the words in the contents
func rankWords(words map[string]int) []models.TopWords {
	var w []models.TopWords
	for k, v := range words {
		w = append(w, models.TopWords{
			Word:  k,
			Count: v,
		})
	}

	sort.Slice(w, func(i, j int) bool {
		return w[i].Count > w[j].Count
	})

	// Get the top 10 most words from the contents
	topWords := make([]models.TopWords, 0, len(w))
	for i := 0; i < 10; i++ {
		topWords = append(topWords, w[i])
	}
	return topWords
}

// countWords counts the number of words in the contents
func countWords(s string) map[string]int {
	words := strings.Fields(s)
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}
	return counts
}

// parseBodyContents parses the JSON payload
func parseBodyContents(body []byte) (*models.Contents, error) {
	errInfo := ""
	// Unmarshall JSON payload into the ContentsPayload struct
	var contents models.ContentsPayload
	err := json.Unmarshal(body, &contents)
	if err != nil {
		errInfo = fmt.Sprintf("error unmarshalling request body: %s", err.Error())
		logger.WithFields(logger.Fields{
			"error": err,
		}).Error(errInfo)
		return nil, err
	}

	// Marshal the project code's struct into a JSON string
	contentData, err := json.Marshal(contents.Contents)
	if err != nil {
		errInfo = fmt.Sprintf("error marshalling contents info: %s", err.Error())
		logger.WithFields(logger.Fields{
			"error": err,
		}).Error(errInfo)
		return nil, err
	}

	// Unmarshall the JSON string into the Contents struct
	var contentDetails models.Contents
	err = json.Unmarshal(contentData, &contentDetails)
	if err != nil {
		errInfo = fmt.Sprintf("error unmarshalling contents info: %s", err.Error())
		logger.WithFields(logger.Fields{
			"error": err,
		}).Error(errInfo)
		return nil, err
	}

	return &contentDetails, nil
}
