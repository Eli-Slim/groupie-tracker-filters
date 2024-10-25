const HiddenLayer = document.getElementsByClassName("container-filter")[0]
const filter = document.getElementById("filter")
const Back = document.getElementById("back")

document.body.addEventListener("click" , (e)=> {
    if (e.target == HiddenLayer) HiddenLayer.style.display = "none"
})

filter.addEventListener("click" , ()=> {
    HiddenLayer.style.display = "flex"
    document.body.classList.add('modal-open')
})

Back.addEventListener("click" , ()=> {
    HiddenLayer.style.display = "none"
})