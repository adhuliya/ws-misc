from django.shortcuts import render
import logging

logger = logging.getLogger("mysite.custom")

# Create your views here.

def index(request):
  context = {}
  return render(request, 'main/index.html', context)


def handler403(request, exception=None, template_name='main/error_page.html'):
  """Display the error page."""
  logger.info("ErrorOccured: %s", exception)
  context = {"statusCode": 403, "message": "Permission Denied"}
  return render(request, 'main/error_page.html', context)


def handler404(request, exception=None, template_name='main/error_page.html'):
  """Display the error page."""
  logger.info("ErrorOccured: %s", exception)
  context = {"statusCode": 404, "message": "Resource Not Found"}
  return render(request, 'main/error_page.html', context)


def handler500(request, exception=None, template_name='main/error_page.html'):
  """Display the error page."""
  logger.info("ErrorOccured: %s", exception)
  context = {"statusCode": 500, "message": "Internal Error"}
  return render(request, 'main/error_page.html', context)


def handler000(request, exception=None, template_name='main/error_page.html'):
  """Handle any error..."""
  logger.info("ErrorOccured: %s", exception)
  context = {"statusCode": "Error", "message": "Error"}
  return render(request, 'main/error_page.html', context)


