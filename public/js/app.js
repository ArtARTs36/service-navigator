function setSearchProvider(url, queryParamName) {
    document.getElementById("search-form").setAttribute("action", url);
    document.getElementById("search-form-query").setAttribute("name", queryParamName);
}

function runSearch() {
    document.querySelector("#search-form").submit();
}

document.addEventListener("DOMContentLoaded", function(event) {
    var popoverTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="popover"]'))
    popoverTriggerList.map(function (popoverTriggerEl) {
        return new bootstrap.Popover(popoverTriggerEl)
    });
});
