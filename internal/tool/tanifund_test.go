package tool_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/sulong/internal/tool"
)

const (
	taniFundURL     = "http://localhost:8080/projects"
	numberOfProject = 10
)

var (
	ctx                   = context.Background()
	invalidTaniFundClient = NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{},
			Body:       ioutil.NopCloser(bytes.NewBufferString(`OK`)),
		}
	})
	validTaniFundClient = NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{},
			Body:       ioutil.NopCloser(bytes.NewBufferString(taniFundProjectJSON())),
		}
	})
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

func TestNewTaniFundCrawler(t *testing.T) {
	t.Run("successfully create TaniFundCrawler", func(t *testing.T) {
		crawler := tool.NewTaniFundCrawler(validTaniFundClient, taniFundURL)
		assert.NotNil(t, crawler)
	})
}

func TestTaniFundCrawler_GetNewestProjects(t *testing.T) {
	t.Run("upstream returns invalid JSON", func(t *testing.T) {
		crawler := tool.NewTaniFundCrawler(invalidTaniFundClient, taniFundURL)
		projects, err := crawler.GetNewestProjects(ctx, numberOfProject)

		assert.NotNil(t, err)
		assert.Empty(t, projects)
	})

	t.Run("upstream returns valid JSON", func(t *testing.T) {
		crawler := tool.NewTaniFundCrawler(validTaniFundClient, taniFundURL)
		projects, err := crawler.GetNewestProjects(ctx, numberOfProject)

		assert.Nil(t, err)
		assert.NotEmpty(t, projects)
	})
}

func taniFundProjectJSON() string {
	return `{"data":{"totalItems":516,"totalPages":86,"currentPage":1,"itemsPerPage":6,"previousLink":null,"nextLink":"https://api.tanifund.com/v2/projects?itemsPerPage=6&page=2&sort=-cutoffAt","items":[{"id":"54f0b448-1abf-4720-aba8-d890f13c762d","projectNo":"P001140521626","farmerGroupId":"57c9e920-cba9-47eb-8148-7493e6dd91a4","title":"Budidaya Nanas Malang - 2","shortDescription":"Pendanaan pada program budidaya nanas dengan bunga 15% p.a","pricePerUnit":100000,"returnPricePerUnit":0,"maxUnit":2554,"startAt":"2021-05-19T17:00:00Z","endAt":"2021-11-19T16:59:00Z","publishedAt":"2021-05-14T12:00:00Z","cutoffAt":"2021-05-19T16:59:00Z","fundraisedAt":"0001-01-01T00:00:00Z","interestPessimist":15,"interestTarget":15,"interestOptimist":15,"grade":"A","youtubeVideoId":"","imageUrl":"https://bucket.tanifund.com/uploads/projects/54f0b448-1abf-4720-aba8-d890f13c762d/120521053708-image.jpg","prospectusUrl":"https://bucket.tanifund.com/uploads/projects/54f0b448-1abf-4720-aba8-d890f13c762d/120521053709-prospectus.pdf","urlSlug":"malang-s-pineapple-cultivation-2","cancelledReason":"","isSupportReferral":1,"isHidden":0,"isSpecialProject":0,"maxUnitPerUser":0,"unitSold":0,"digitalSignatureDocumentId":null,"digitalSignatureDocumentSignedAt":null,"digitalSignatureDocumentUrl":null,"projectStatusOrder":2,"projectStatus":{"id":5,"description":"Menunggu Fundraising","order":2},"fundingType":{"id":2,"description":"Pengembalian tetap"},"returnPeriod":{"id":2,"description":"Bulanan"},"cultivation":{"id":67,"description":"Nanas","isActive":1,"cultivationCategory":{"id":1,"description":"Pertanian"}},"village":{"id":"3507170005","description":"Senggreng","district":{"id":"3507170","description":"Sumber Pucung","city":{"id":"3507","description":"Kab. Malang","province":{"id":"35","description":"Jawa Timur"}}}}},{"id":"353e81d1-3d47-406a-8ff2-595b4a059f13","projectNo":"P001130521625","farmerGroupId":"e582dd6b-ab12-4d44-bba2-13c7e7c83375","title":"Budidaya Kabocha Malang - 5","shortDescription":"Pendanaan pada program budidaya kabocha dengan bunga 15% p.a","pricePerUnit":100000,"returnPricePerUnit":0,"maxUnit":1557,"startAt":"2021-05-18T17:00:00Z","endAt":"2021-09-18T16:59:00Z","publishedAt":"2021-05-13T12:00:00Z","cutoffAt":"2021-05-18T16:59:00Z","fundraisedAt":"0001-01-01T00:00:00Z","interestPessimist":15,"interestTarget":15,"interestOptimist":15,"grade":"A","youtubeVideoId":"","imageUrl":"https://bucket.tanifund.com/uploads/projects/353e81d1-3d47-406a-8ff2-595b4a059f13/120521051636-image.jpg","prospectusUrl":"https://bucket.tanifund.com/uploads/projects/353e81d1-3d47-406a-8ff2-595b4a059f13/120521051636-prospectus.pdf","urlSlug":"malang-s-kabocha-cultivation-5","cancelledReason":"","isSupportReferral":1,"isHidden":0,"isSpecialProject":0,"maxUnitPerUser":0,"unitSold":0,"digitalSignatureDocumentId":null,"digitalSignatureDocumentSignedAt":null,"digitalSignatureDocumentUrl":null,"projectStatusOrder":2,"projectStatus":{"id":5,"description":"Menunggu Fundraising","order":2},"fundingType":{"id":2,"description":"Pengembalian tetap"},"returnPeriod":{"id":2,"description":"Bulanan"},"cultivation":{"id":81,"description":"Labu","isActive":1,"cultivationCategory":{"id":1,"description":"Pertanian"}},"village":{"id":"3507160013","description":"Ardirejo","district":{"id":"3507160","description":"Kepanjen","city":{"id":"3507","description":"Kab. Malang","province":{"id":"35","description":"Jawa Timur"}}}}},{"id":"9fad1295-8b55-4f26-a252-dba5c90ed566","projectNo":"P003120521597","farmerGroupId":"6a7903b4-a113-4209-b4de-a5826ee41c01","title":"Budidaya Ikan Nila Tasikmalaya - 1","shortDescription":"Pendanaan pada program budidaya ikan nila dengan bunga 17% p.a","pricePerUnit":100000,"returnPricePerUnit":0,"maxUnit":9651,"startAt":"2021-05-19T17:00:00Z","endAt":"2021-11-19T16:59:00Z","publishedAt":"2021-05-12T03:00:00Z","cutoffAt":"2021-05-19T16:59:00Z","fundraisedAt":"2021-05-12T04:30:57Z","interestPessimist":17,"interestTarget":17,"interestOptimist":17,"grade":"A","youtubeVideoId":"","imageUrl":"https://bucket.tanifund.com/uploads/projects/9fad1295-8b55-4f26-a252-dba5c90ed566/110521081506-image.jpg","prospectusUrl":"https://bucket.tanifund.com/uploads/projects/9fad1295-8b55-4f26-a252-dba5c90ed566/110521103058-prospectus.pdf","urlSlug":"tasikmalaya-nile-tilapia-fish-cultivation-1","cancelledReason":"","isSupportReferral":1,"isHidden":0,"isSpecialProject":0,"maxUnitPerUser":0,"unitSold":9651,"digitalSignatureDocumentId":null,"digitalSignatureDocumentSignedAt":null,"digitalSignatureDocumentUrl":null,"projectStatusOrder":3,"projectStatus":{"id":1,"description":"Menunggu Dimulai","order":3},"fundingType":{"id":2,"description":"Pengembalian tetap"},"returnPeriod":{"id":2,"description":"Bulanan"},"cultivation":{"id":74,"description":"Ikan Nila","isActive":1,"cultivationCategory":{"id":3,"description":"Perikanan"}},"village":{"id":"3206260003","description":"Bugel","district":{"id":"3206260","description":"Ciawi","city":{"id":"3206","description":"Kab. Tasikmalaya","province":{"id":"32","description":"Jawa Barat"}}}}},{"id":"b24267ec-e2b9-4bc8-b7ef-550bf8394595","projectNo":"P003110521578","farmerGroupId":"a60a08fa-2665-4a83-8ef4-a0e27b0f91c4","title":"Budidaya Ikan Mas Subang - 1","shortDescription":"Pendanaan pada program budidaya ikan mas dengan bunga 15% p.a","pricePerUnit":100000,"returnPricePerUnit":0,"maxUnit":822,"startAt":"2021-05-13T17:00:00Z","endAt":"2021-10-13T16:59:00Z","publishedAt":"2021-05-11T08:00:00Z","cutoffAt":"2021-05-13T16:59:00Z","fundraisedAt":"2021-05-11T11:52:38Z","interestPessimist":15,"interestTarget":15,"interestOptimist":15,"grade":"A","youtubeVideoId":"","imageUrl":"https://bucket.tanifund.com/uploads/projects/b24267ec-e2b9-4bc8-b7ef-550bf8394595/110521080335-image.jpg","prospectusUrl":"https://bucket.tanifund.com/uploads/projects/b24267ec-e2b9-4bc8-b7ef-550bf8394595/110521102601-prospectus.pdf","urlSlug":"subang-s-goldfish-cultivation-1","cancelledReason":"","isSupportReferral":1,"isHidden":0,"isSpecialProject":0,"maxUnitPerUser":0,"unitSold":822,"digitalSignatureDocumentId":null,"digitalSignatureDocumentSignedAt":null,"digitalSignatureDocumentUrl":null,"projectStatusOrder":3,"projectStatus":{"id":1,"description":"Menunggu Dimulai","order":3},"fundingType":{"id":2,"description":"Pengembalian tetap"},"returnPeriod":{"id":2,"description":"Bulanan"},"cultivation":{"id":94,"description":"Ikan Mas","isActive":1,"cultivationCategory":{"id":3,"description":"Perikanan"}},"village":{"id":"3213040001","description":"Buniara","district":{"id":"3213040","description":"Tanjungsiang","city":{"id":"3213","description":"Kab. Subang","province":{"id":"32","description":"Jawa Barat"}}}}},{"id":"21c91fbb-3a71-4c76-86f9-ce67c54669db","projectNo":"P001100521620","farmerGroupId":"e5853232-4061-4f2a-9104-8385d73dcdc7","title":"Budidaya Jagung Pipil Lombok - 2","shortDescription":"Pendanaan pada program budidaya jagung pipil dengan bunga 11% p.a","pricePerUnit":100000,"returnPricePerUnit":0,"maxUnit":10000,"startAt":"2021-05-16T17:00:00Z","endAt":"2021-08-15T16:59:00Z","publishedAt":"2021-05-10T12:00:00Z","cutoffAt":"2021-05-16T16:59:00Z","fundraisedAt":"2021-05-10T22:40:47Z","interestPessimist":11,"interestTarget":11,"interestOptimist":11,"grade":"A","youtubeVideoId":"","imageUrl":"https://bucket.tanifund.com/uploads/projects/21c91fbb-3a71-4c76-86f9-ce67c54669db/100521090220-image.jpg","prospectusUrl":"https://bucket.tanifund.com/uploads/projects/21c91fbb-3a71-4c76-86f9-ce67c54669db/100521090220-prospectus.pdf","urlSlug":"lombok-s-shelled-corn-cultivation-2","cancelledReason":"","isSupportReferral":1,"isHidden":0,"isSpecialProject":0,"maxUnitPerUser":0,"unitSold":10000,"digitalSignatureDocumentId":null,"digitalSignatureDocumentSignedAt":null,"digitalSignatureDocumentUrl":null,"projectStatusOrder":3,"projectStatus":{"id":1,"description":"Menunggu Dimulai","order":3},"fundingType":{"id":2,"description":"Pengembalian tetap"},"returnPeriod":{"id":3,"description":"Akhir Periode"},"cultivation":{"id":32,"description":"Jagung Pipil","isActive":1,"cultivationCategory":{"id":1,"description":"Pertanian"}},"village":{"id":"5203090016","description":"Lenek Pesiraman","district":{"id":"5203090","description":"Aikmel","city":{"id":"5203","description":"Kab. Lombok Timur","province":{"id":"52","description":"Nusa Tenggara Barat"}}}}},{"id":"6bd702c4-6198-4d98-9d60-b5b3ceb8e4be","projectNo":"P001080521623","farmerGroupId":"aa82a9fd-e3b3-4320-9820-a406d3f59092","title":"Pengembangan Usaha Restoran Burger KB - 6","shortDescription":"Pendanaan pada program pembiayaan usaha restauran burger dengan bagi hasil 15% p.a","pricePerUnit":100000,"returnPricePerUnit":0,"maxUnit":9424,"startAt":"2021-05-19T17:00:00Z","endAt":"2021-08-18T16:59:00Z","publishedAt":"2021-05-08T15:00:00Z","cutoffAt":"2021-05-19T16:59:00Z","fundraisedAt":"2021-05-11T01:23:11Z","interestPessimist":15,"interestTarget":15,"interestOptimist":15,"grade":"A","youtubeVideoId":"","imageUrl":"https://bucket.tanifund.com/uploads/projects/6bd702c4-6198-4d98-9d60-b5b3ceb8e4be/080521135354-image.jpg","prospectusUrl":"https://bucket.tanifund.com/uploads/projects/6bd702c4-6198-4d98-9d60-b5b3ceb8e4be/080521135355-prospectus.pdf","urlSlug":"kb-burger-restaurant-business-financing-6","cancelledReason":"","isSupportReferral":1,"isHidden":0,"isSpecialProject":0,"maxUnitPerUser":0,"unitSold":9424,"digitalSignatureDocumentId":null,"digitalSignatureDocumentSignedAt":null,"digitalSignatureDocumentUrl":null,"projectStatusOrder":3,"projectStatus":{"id":1,"description":"Menunggu Dimulai","order":3},"fundingType":{"id":2,"description":"Pengembalian tetap"},"returnPeriod":{"id":2,"description":"Bulanan"},"cultivation":{"id":48,"description":"Aneka Bahan Pangan","isActive":1,"cultivationCategory":{"id":4,"description":"Kompilasi"}},"village":{"id":"3172090007","description":"Pulo Gadung","district":{"id":"3172090","description":"Pulo Gadung","city":{"id":"3172","description":"Kota Jakarta Timur","province":{"id":"31","description":"DKI Jakarta"}}}}}]}}`
}
