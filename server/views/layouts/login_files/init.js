seajs.config({
  paths: {
    'T': 'templates',
    'C': 'controllers',
    'CM': 'components',
    'css': '../css',
  },
  alias: {
    'app': 'app',
    'tpl': 'template',
    'ui': 'libs/jquery-ui.min',
    'underscore': 'libs/underscore-min',
    'backbone': 'libs/backbone-min',
    'utils': 'libs/utils.js',
    'moment': 'libs/moment.min',
    'template': 'libs/seajs-text-debug',
    'store': 'libs/store.js',
  },
  map: [
    [/^([^.]+)$/i, '$1.js='+staticVersion],
    [/^(.+\.(?:html|js|css))$/i, '$1?v='+staticVersion]
  ],
  base: staticUri + '/pc/js/',
  charset: 'utf-8'
});

if (typeof customApp !== 'undefined' || !_.isEmpty(customApp)) { 
    jQuery(function(){
        seajs.use(['underscore', 'store', 'utils', 'template', 'ui'], function() {
            seajs.use(customApp[0], function(loader) {
                if (!_.isNull(loader)) {
                    window.Main = new loader[ customApp[1] ](customApp[2]);
                }
            });
        });
    });
}

// Banner Init
seajs.use("C/item/menu");
// wxLogin
// 谁将此代码放这里的？请将其移动到Login页面中再触发.
//seajs.use("C/sign/wxLogin");
