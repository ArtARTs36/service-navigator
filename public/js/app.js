function setSearchProvider(url, queryParamName) {
    document.getElementById("search-form").setAttribute("action", url);
    document.getElementById("search-form-query").setAttribute("name", queryParamName);
}

function runSearch() {
    document.querySelector("#search-form").submit();
}
