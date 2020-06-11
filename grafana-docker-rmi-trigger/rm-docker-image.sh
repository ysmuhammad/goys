test=$(date +"%Y-%m-%d")T$(date -d '2 hours ago' +"%T")
docker image prune -a --force --filter "until=${test}"