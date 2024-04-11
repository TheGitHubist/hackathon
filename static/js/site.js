function openCategory(category) {
    var i;
    var x = document.getElementsByClassName("tab"); // Utilisation de getElementsByClassName au lieu de getElementById
    for (i = 0; i < x.length; i++) {
        x[i].style.backgroundColor = "#d3d3d3";
    }
    var x = document.getElementsByClassName("mainContent"); // Utilisation de getElementsByClassName au lieu de getElementById
    for (i = 0; i < x.length; i++) {
        x[i].style.display = "none";
    }
    document.getElementById(category).style.backgroundColor = "#fdc500"
    document.getElementById(category).style.display = "block"; // Correction ici : getElementById au lieu de getElementBy

    document.getElementById(category + "-container").style.display = "block"
}


function redirectToPost(id) {

    openCategory(id);
    changeColours()
}


function changeColours() {
    var cols = document.querySelectorAll('.subnavbar .col');
    alert(col)
}