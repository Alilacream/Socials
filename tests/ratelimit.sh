for i in {1..11}; do
  curl http://localhost:8080/health/ok
  echo "request $i is done "
done
