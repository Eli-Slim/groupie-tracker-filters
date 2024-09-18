const API_URL = "http://localhost:8080/";
const GEOCODE_API_URL = "https://nominatim.openstreetmap.org/search.php?q="

let id = document.querySelector(".team-content").dataset.id;


async function initMap() {
  const loaderEl = loader(".map-content")
  let map = L.map("map").setView([51.505, -0.09], 2);
  L.tileLayer("https://tile.openstreetmap.org/{z}/{x}/{y}.png", {
    maxZoom: 19,
    attribution:
      '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>',
  }).addTo(map);
  await SetLocationsOnMap(map);
  loaderEl.remove()
}

initMap()


