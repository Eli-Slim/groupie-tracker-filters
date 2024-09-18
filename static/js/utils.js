
async function fetchJson(url, options = {}) {
    const response = await fetch(url, options)
    if (!response.ok) {
      throw new Error("HTTP error! status: ${response.status}")
    }
    const data = await response.json()
    return data
  }
  
  async function getGeoCode(city, country) {
    const url = encodeURI(`${GEOCODE_API_URL}+${city},${country}&format=jsonv2`)
    const data = await fetchJson(url)
    if (data.length) {
      let coordinate = data[0]
      return {lat: coordinate.lat, lon: coordinate.lon}
    }
    throw new Error("Location not found");
  }
  
  async function getLocation(id) {
    const url = `${API_URL}locations/${id}`;
    const data = await fetchJson(url)
    return data
  }
  
  function formatLocation(location) {
    let [city, country] = location.split("-");
    city = city.replaceAll("_", " ");
    return [city, country]
  }
  
  async function getLocalizationByID(id) {
    try {
      const res = [];
      const locations = await getLocation(id)
      for (let location of locations.locations) {
        let [city, country] = formatLocation(location)
        let coordinate = await getGeoCode(city, country);
        res.push({ ...coordinate, city, country });
      }
      return res;
    } catch (error) {
      console.log(error);
    }
  }
  
  async function SetLocationsOnMap(map) {
    let coordinates = await getLocalizationByID(id);
    for (let coordinate of coordinates) {
      let { lat, lon, city, country } = coordinate;
      L.marker([lat, lon])
        .addTo(map)
        .bindPopup(`${city}, ${country}`)
    }
  }


function loader(parent) {
  let parentEl = document.querySelector(parent)
  const html = `
  <div class="loading">
    <div class="loader"></div>
    <span>Loading...</span>
  </div>`
  parentEl.insertAdjacentHTML("beforeend", html)
  return parentEl.querySelector(".loading")
}