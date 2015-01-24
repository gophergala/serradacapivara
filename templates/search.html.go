<!DOCTYPE html>
<html lang="pt-br">
    <head>
        <title>Serra da Capivara</title>
        <meta charset="utf-8" />
        <link type="text/css" rel="stylesheet" href="/static/css/layout.css"  media="screen,projection"/>
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no"/>
        <script src="https://maps.googleapis.com/maps/api/js?v=3.exp&signed_in=true"></script>
        <script>
            var map;
            function initialize() {
                var mapOptions = {
                    scrollwheel: false,
                    zoom: 8,
                    center: new google.maps.LatLng(-34.397, 150.644)
                };
                map = new google.maps.Map(document.getElementById('map-canvas'),
                    mapOptions);
            }
            
            google.maps.event.addDomListener(window, 'load', initialize);
            
        </script>
    </head>
    <body class="index">
        <div class="navbar-fixed">
            <nav class="transparent" role="navigation">
                <div class="container">
                    <div class="nav-wrapper">
                        <a id="logo-container" href="#" class="brand-logo thin-text">
                        Resultados para "texto"
                        </a>
                        <ul class="side-nav">
                            <li><a href="#"><i class="mdi-action-search"></i></a></li>
                            <li><a href="#"  data-activates="slide-out" class="button-menu"><i class="mdi-navigation-more-vert"></i></a></li>
                        </ul>
                        <ul id="slide-out" class="side-nav full">
                            <li><a href="#">ABOUT US</a></li>
                            <li><a href="#">HELP US</a></li>
                            <li><a href="#">SEARCH</a></li>
                            <li><a href="#">FAQ</a></li>
                        </ul>
                        <a href="javascript:history.back();">
                        <i class="mdi-navigation-arrow-back"></i>
                        </a>
                    </div>
                </div>
            </nav>
        </div>
        <div class="map-results">
            <div class="container">
                <div class="row">
                    <div class="col s12 m6">
                        <div class="card">
                            <div class="card-content">
                                <span class="card-title grey-text">Results</span>
                                <p>Results founded for "texto"</p>
                                <div class="collection">
                                    <a href="#!" class="collection-item">Sitio 1</a>
                                    <a href="#!" class="collection-item active">Sitio 8</a>
                                    <a href="#!" class="collection-item">Sitio 2</a>
                                    <a href="#!" class="collection-item">Sitio 14</a>
                                </div>
                            </div>
                            <div class="card-action">
                                <a href="#">Close</a>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div id="map-canvas" class="full"></div>
        <footer class="page-footer">
            <div class="footer-copyright">
                <div class="container">
                    Made with Go, Love, Doritos and <a class="deep-orange-text darken-4" href="http://materializecss.com">Materialize</a>
                </div>
            </div>
        </footer>
        <!--Import jQuery before materialize.js-->
        <script type="text/javascript" src="/static/js/jquery.js"></script>
        <script type="text/javascript" src="/static/js/materialize.min.js"></script>
        <script type="text/javascript">
            $(document).ready(function(){
                $('.parallax').parallax();
                $('.button-menu').sideNav();
            });
        </script>
    </body>
</html>