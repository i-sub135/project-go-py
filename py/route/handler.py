from app import app
from flask import render_template, request, redirect
import requests

from setting import BASE_API




@app.route("/", methods=["GET"])
def index():
    return "index"


@app.route('/hello', methods = ["GET"])
def hello():
   return render_template("child/hello.html")

@app.route('/DataCarfix', methods = ["GET"])
def DataCarfix():
   

   if not request.args.get('pages') or int(request.args.get('pages')) < 0:
       return redirect("/DataCarfix?pages=0")
       

   pages =  int(request.args.get('pages')) >= 0 and int(request.args.get('pages')) or 0
   
   
   halaman = 10
   page = pages > 1 and (pages * halaman) - halaman or 0


   uri = BASE_API + "/api/v1/carfix?pages=" + str(halaman) + "&mulai=" + str(page)
   data = requests.get(url = uri)

   num = 0
   d = data.json()['result']
   return render_template("child/carfix.html", data = d, page = pages , num = num)