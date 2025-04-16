!! Create docker image by following command 
docker build -t date_and_time:v1 .

!! Create docker Container by following command 
docker run -it --rm -p 8080:8080 --name pruthvi_con date_and_time:v1

!! To scan the images using kubescape 
docker run --rm \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v $PWD:/output \
  bitnami/kubescape \
  scan image date_and_time:v1 \
  --format json \
  --output /output/data_and_time_image_scan_findings.json

  !! Check web application ( data and time ) at following url
  https://product-management-usu2.onrender.com/
