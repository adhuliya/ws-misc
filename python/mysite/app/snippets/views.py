from django.shortcuts import render
from django.http import HttpResponse, Http404, HttpResponseRedirect, JsonResponse
from django.utils import timezone
from django.urls import reverse
import json
import logging

logger = logging.getLogger("mysite.custom")

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
  count = request.session.get("visitCount", 0)
  count += 1
  request.session["visitCount"] = count

  # # Set session as modified to force data updates/cookie to be saved.
  # request.session.modified = True

  return render(request, "snippets/visit_count.html", {"visitCount": count})


def protectFile(request, filepath):
  """Serves static files that are protected,
  and uses nginx to server the files very efficiently.
  Ref: https://wellfire.co/learn/nginx-django-x-accel-redirects/
  """
  logger.info("ProtectFile: %s", filepath)
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
  """Sample view to receive json data and send the same back
  to client."""
  if request.is_ajax():
    if request.method == 'POST':
      jsonData = json.loads(request.body)
      try:
        data = jsonData['data']
      except KeyError:
        Http404("Malformed data! (expected json)")
      else:
        return JsonResponse(data)
  return HttpResponse("No json data received.")


def displayRequest(request):
  """Displays the information received in the request."""
  context = {"request": request}
  return render(request, "snippets/display_request.html", context)


def loggingDemo(request):
  """A simple view to demonstrate the logging system."""
  import logging
  # get the mysite.custom logger defined in mysite.settings module
  logger = logging.getLogger("mysite.custom")
  message = "LoggingDemoSuccessful"
  logger.info(message)
  return render(request, "snippets/logging_demo.html", {"message":message})


def raiseError(request, statusCode=404):
  """Raise the error corresponding to the given status code."""
  from django.core.exceptions import PermissionDenied
  if statusCode == 403:
    raise PermissionDenied    # raises 403
  if statusCode == 404:
    raise Http404             # raises 404
  if statusCode == 500:
    raise Exception()         # raises 500
  else:
    # For all others redirect to the index page
    return HttpResponseRedirect(reverse('main:index'))


