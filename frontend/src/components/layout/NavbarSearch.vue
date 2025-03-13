<template>
  {% if _navBar.Search.Providers | length > 0 %}
  <form class="col-12 col-lg-auto mb-3 mb-lg-0 me-lg-3" action="{{ _navBar.Search.Providers[0].URL }}" id="search-form">
    <input
        type="search"
        class="form-control"
        name="{{ _navBar.Search.Providers[0].QueryParamName }}"
        placeholder="Search..."
        aria-label="Search"
        id="search-form-query"
    />
  </form>

  {% if _navBar.Search.Providers | length == 1 %}
  <a href="#" class="d-block link-dark text-decoration-none" onclick="runSearch()">
    <i class="fa-solid fa-search"></i>
  </a>
  {% else %}
  <div class="dropdown">
    <a href="#" class="d-block link-dark text-decoration-none dropdown-toggle" id="dropdownUser2" data-bs-toggle="dropdown" aria-expanded="false">
      <i class="fa-solid fa-search"></i>
    </a>
    <ul class="dropdown-menu text-small" aria-labelledby="dropdownUser2">
      {% for provider in _navBar.Search.Providers %}
      <li><a class="dropdown-item search-dropdown-item" onclick="setSearchProvider('{{ provider.URL }}', '{{ provider.QueryParamName }}'); runSearch()">{{ provider.Name }}</a></li>
      {% endfor %}
    </ul>
  </div>
  {% endif %}
  {% endif %}

  <div>&nbsp;</div>
  <div>&nbsp;</div>
</template>
