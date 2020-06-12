from django.shortcuts import render
from django.http import HttpResponse, Http404, HttpResponseRedirect, JsonResponse
from django.utils import timezone
import json

# Create your views here.

def index(request):
  return render(request, "snippets/index.html", {})


def visitCount(request):
  """
  Count the visits by the client browser and show it.
  It demonstrates using the session cookies.
  Ref: https://developer.mozilla.org/en-US/docs/Learn/Server-side/Django/Sessions
  """
  # Number of visits to this view, as counted in the session variable.
  visitCount = request.session.get("visitCount", 0)
  visitCount += 1
  request.session["visitCount"] = visitCount

  # # Set session as modified to force data updates/cookie to be saved.
  # request.session.modified = True

  return render(request, "snippets/visit_count.html", {"visitCount": visitCount})


def protectFile(request, filepath):
  """Serves static files that are protected,
  and uses nginx to server the files very efficiently.
  Ref: https://wellfire.co/learn/nginx-django-x-accel-redirects/
  """
  print("ProtectFile:", filepath) #delit
  response = HttpResponse()
  response["Content-Disposition"] = "attachment; filename={0}".format(filepath)
  response['X-Accel-Redirect'] = f"/protected/{filepath}"
  return response


def jsonSend(request):
  """Return json response."""
  data = {
    'time': str(timezone.now()),
    'is_active': True,
    'count': 28
  }
  return JsonResponse(data)


def jsonReceive(request):
  """Sample view to receive json data."""
  if request.is_ajax():
    if request.method == 'POST':
      jsonData = json.loads(request.body)
      try:
        data = jsonData['data']
      except KeyError:
        Http404("Malformed data!")
      return HttpResponse("Got json data")
  return HttpResponse("No json data.")


