# Anastasiia go-pets
Prerequisites:
Need to create a Service Account with access rights - Owner, Datastore Editor
Create and download a service account key to your local machine
Enter your project ID and path to your service account key whenever promoted

To run execute: ./start_my_app.sh To destroy execute: ./stop_my_app.sh

For POST API that adds new record to DB, run CURL command in command line (update the field values if you want):

curl -X POST -H "Content-Type: application/json" -d '{"Added":"2021-09-30T02:02:04.373Z","Caption":"biker duck","Email":"duck@gamil.com","Image":"https://www.goodnewsnetwork.org/wp-content/uploads/2015/08/guus-the-duck-motorcycle-Instagram.jpg","Likes":250,"Owner":"Greg","Petname":"Devil"}' https://nastia-go-pets-4xlapmsusq-uc.a.run.app/pets
