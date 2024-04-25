from django.contrib import admin
from django.urls import path, include

from playground.views import index, contact


# hello
urlpatterns = [
    path("", index, name="index"),
    path("contact/", contact, name="contact"),
    path("admin/", admin.site.urls),
]
