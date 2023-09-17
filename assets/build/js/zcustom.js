

document.addEventListener('DOMContentLoaded', function () {
    var elems = document.querySelectorAll('.sidenav');
    var instances = M.Sidenav.init(elems, options);
    var instance = M.Tabs.init(el, options);

});

$(document).ready(function () {
    $('.sidenav').sidenav();
    $('.tabs').tabs();
    var tooltipTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="tooltip"]'));
    var tooltipList = tooltipTriggerList.map(function (tooltipTriggerEl) {
        return new bootstrap.Tooltip(tooltipTriggerEl);
    });
    $('tabs', this).click(function () {

    });

});

var instance = M.Tabs.init(el, options);

// Or with jQuery

$(document).ready(function(){
  $('.tabs').tabs();
});