FROM registry.ci.openshift.org/ocp/builder:rhel-8-golang-1.19-openshift-4.12 AS builder
WORKDIR /go/src/github.com/openshift/route-controller-manager
COPY . .
RUN make build --warn-undefined-variables

FROM registry.ci.openshift.org/ocp/4.12:base
COPY --from=builder /go/src/github.com/openshift/route-controller-manager/test /usr/bin/
LABEL io.k8s.display-name="OpenShift Route Controller Manager Command" \
      io.k8s.description="OpenShift is a platform for developing, building, and deploying containerized applications." \
      io.openshift.tags="openshift,route-controller-manager"
