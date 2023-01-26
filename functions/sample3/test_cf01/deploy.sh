ENTRY_POINT="HelloWorld"
REGION="asia-northeast1"
PROJECT=""
gcloud functions deploy test02 \
--project ${PROJECT} \
--gen2 \
--runtime=go119 \
--region=${REGION} \
--source=function/. \
--memory=128Mi \
--entry-point=${ENTRY_POINT} \
--timeout 600s \
--trigger-http \
--allow-unauthenticated