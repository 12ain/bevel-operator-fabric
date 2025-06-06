package hlfmetrics

import (
	"crypto/x509"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	CertificateExpiryTimeSeconds = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "hlf_operator_certificate_expiration_timestamp_seconds",
			Help: "The date after which the certificate expires. Expressed as a Unix Epoch Time.",
		},
		[]string{"node_type", "crt_type", "namespace", "name"},
	)

	CurrentTimeSeconds = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "hlf_operator_current_time_seconds",
			Help: "The current time in Unix Epoch Time.",
		},
	)
)

func UpdateCertificateExpiry(
	nodeType string,
	crtType string,
	crt *x509.Certificate,
	name string,
	ns string,
) {
	expiryTime := 0.0
	expiryTime = float64(crt.NotAfter.Unix())
	CertificateExpiryTimeSeconds.With(prometheus.Labels{
		"namespace": ns,
		"name":      name,
		"node_type": nodeType,
		"crt_type":  crtType,
	}).Set(expiryTime)
}

func UpdateCurrentTime() {
	CurrentTimeSeconds.Set(float64(time.Now().Unix()))
}
