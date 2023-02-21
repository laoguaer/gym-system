import{i as e,o as t}from"./utils.3e78d745.js";import{by as o,U as n,j as i,m as r,p as l,a3 as a,az as s,i as u,aK as c,V as d,k as h,bz as f,b1 as v,a7 as p,Z as g,z as w,a4 as m,an as C,ap as b,aT as x,Q as P,q as O,n as S,x as L,bn as z,bA as I,br as k,bB as T,T as y,ak as M,N as R,bs as A,C as E,aP as D,b0 as H,as as Z,af as j,ad as $,ac as B,ae as V}from"./index.9c865cdf.js";import{u as N}from"./Input.463ffced.js";var W;const U=(W={"À":"A","Á":"A","Â":"A","Ã":"A","Ä":"A","Å":"A","à":"a","á":"a","â":"a","ã":"a","ä":"a","å":"a","Ç":"C","ç":"c","Ð":"D","ð":"d","È":"E","É":"E","Ê":"E","Ë":"E","è":"e","é":"e","ê":"e","ë":"e","Ì":"I","Í":"I","Î":"I","Ï":"I","ì":"i","í":"i","î":"i","ï":"i","Ñ":"N","ñ":"n","Ò":"O","Ó":"O","Ô":"O","Õ":"O","Ö":"O","Ø":"O","ò":"o","ó":"o","ô":"o","õ":"o","ö":"o","ø":"o","Ù":"U","Ú":"U","Û":"U","Ü":"U","ù":"u","ú":"u","û":"u","ü":"u","Ý":"Y","ý":"y","ÿ":"y","Æ":"Ae","æ":"ae","Þ":"Th","þ":"th","ß":"ss","Ā":"A","Ă":"A","Ą":"A","ā":"a","ă":"a","ą":"a","Ć":"C","Ĉ":"C","Ċ":"C","Č":"C","ć":"c","ĉ":"c","ċ":"c","č":"c","Ď":"D","Đ":"D","ď":"d","đ":"d","Ē":"E","Ĕ":"E","Ė":"E","Ę":"E","Ě":"E","ē":"e","ĕ":"e","ė":"e","ę":"e","ě":"e","Ĝ":"G","Ğ":"G","Ġ":"G","Ģ":"G","ĝ":"g","ğ":"g","ġ":"g","ģ":"g","Ĥ":"H","Ħ":"H","ĥ":"h","ħ":"h","Ĩ":"I","Ī":"I","Ĭ":"I","Į":"I","İ":"I","ĩ":"i","ī":"i","ĭ":"i","į":"i","ı":"i","Ĵ":"J","ĵ":"j","Ķ":"K","ķ":"k","ĸ":"k","Ĺ":"L","Ļ":"L","Ľ":"L","Ŀ":"L","Ł":"L","ĺ":"l","ļ":"l","ľ":"l","ŀ":"l","ł":"l","Ń":"N","Ņ":"N","Ň":"N","Ŋ":"N","ń":"n","ņ":"n","ň":"n","ŋ":"n","Ō":"O","Ŏ":"O","Ő":"O","ō":"o","ŏ":"o","ő":"o","Ŕ":"R","Ŗ":"R","Ř":"R","ŕ":"r","ŗ":"r","ř":"r","Ś":"S","Ŝ":"S","Ş":"S","Š":"S","ś":"s","ŝ":"s","ş":"s","š":"s","Ţ":"T","Ť":"T","Ŧ":"T","ţ":"t","ť":"t","ŧ":"t","Ũ":"U","Ū":"U","Ŭ":"U","Ů":"U","Ű":"U","Ų":"U","ũ":"u","ū":"u","ŭ":"u","ů":"u","ű":"u","ų":"u","Ŵ":"W","ŵ":"w","Ŷ":"Y","ŷ":"y","Ÿ":"Y","Ź":"Z","Ż":"Z","Ž":"Z","ź":"z","ż":"z","ž":"z","Ĳ":"IJ","ĳ":"ij","Œ":"Oe","œ":"oe","ŉ":"'n","ſ":"s"},function(e){return null==W?void 0:W[e]});var Y=/[\xc0-\xd6\xd8-\xf6\xf8-\xff\u0100-\u017f]/g,X=RegExp("[\\u0300-\\u036f\\ufe20-\\ufe2f\\u20d0-\\u20ff]","g");var F=/[^\x00-\x2f\x3a-\x40\x5b-\x60\x7b-\x7f]+/g;var G=/[a-z][A-Z]|[A-Z]{2}[a-z]|[0-9][a-zA-Z]|[a-zA-Z][0-9]|[^a-zA-Z0-9 ]/;var q="\\ud800-\\udfff",J="\\u2700-\\u27bf",K="a-z\\xdf-\\xf6\\xf8-\\xff",_="A-Z\\xc0-\\xd6\\xd8-\\xde",Q="\\xac\\xb1\\xd7\\xf7\\x00-\\x2f\\x3a-\\x40\\x5b-\\x60\\x7b-\\xbf\\u2000-\\u206f \\t\\x0b\\f\\xa0\\ufeff\\n\\r\\u2028\\u2029\\u1680\\u180e\\u2000\\u2001\\u2002\\u2003\\u2004\\u2005\\u2006\\u2007\\u2008\\u2009\\u200a\\u202f\\u205f\\u3000",ee="["+Q+"]",te="\\d+",oe="["+J+"]",ne="["+K+"]",ie="[^"+q+Q+te+J+K+_+"]",re="(?:\\ud83c[\\udde6-\\uddff]){2}",le="[\\ud800-\\udbff][\\udc00-\\udfff]",ae="["+_+"]",se="(?:"+ne+"|"+ie+")",ue="(?:"+ae+"|"+ie+")",ce="(?:['’](?:d|ll|m|re|s|t|ve))?",de="(?:['’](?:D|LL|M|RE|S|T|VE))?",he="(?:[\\u0300-\\u036f\\ufe20-\\ufe2f\\u20d0-\\u20ff]|\\ud83c[\\udffb-\\udfff])?",fe="[\\ufe0e\\ufe0f]?",ve=fe+he+("(?:\\u200d(?:"+["[^"+q+"]",re,le].join("|")+")"+fe+he+")*"),pe="(?:"+[oe,re,le].join("|")+")"+ve,ge=RegExp([ae+"?"+ne+"+"+ce+"(?="+[ee,ae,"$"].join("|")+")",ue+"+"+de+"(?="+[ee,ae+se,"$"].join("|")+")",ae+"?"+se+"+"+ce,ae+"+"+de,"\\d*(?:1ST|2ND|3RD|(?![123])\\dTH)(?=\\b|[a-z_])","\\d*(?:1st|2nd|3rd|(?![123])\\dth)(?=\\b|[A-Z_])",te,pe].join("|"),"g");function we(e,t,n){return e=o(e),void 0===(t=n?void 0:t)?function(e){return G.test(e)}(e)?function(e){return e.match(ge)||[]}(e):function(e){return e.match(F)||[]}(e):e.match(t)||[]}var me,Ce=RegExp("['’]","g");const be=(me=function(e,t,o){return e+(o?"-":"")+t.toLowerCase()},function(e){return function(e,t,o,n){var i=-1,r=null==e?0:e.length;for(n&&r&&(o=e[++i]);++i<r;)o=t(o,e[i],i,e);return o}(we(function(e){return(e=o(e))&&e.replace(Y,U).replace(X,"")}(e).replace(Ce,"")),me,"")}),xe=n("rotateClockwise",i("svg",{viewBox:"0 0 20 20",fill:"none",xmlns:"http://www.w3.org/2000/svg"},i("path",{d:"M3 10C3 6.13401 6.13401 3 10 3C13.866 3 17 6.13401 17 10C17 12.7916 15.3658 15.2026 13 16.3265V14.5C13 14.2239 12.7761 14 12.5 14C12.2239 14 12 14.2239 12 14.5V17.5C12 17.7761 12.2239 18 12.5 18H15.5C15.7761 18 16 17.7761 16 17.5C16 17.2239 15.7761 17 15.5 17H13.8758C16.3346 15.6357 18 13.0128 18 10C18 5.58172 14.4183 2 10 2C5.58172 2 2 5.58172 2 10C2 10.2761 2.22386 10.5 2.5 10.5C2.77614 10.5 3 10.2761 3 10Z",fill:"currentColor"}),i("path",{d:"M10 12C11.1046 12 12 11.1046 12 10C12 8.89543 11.1046 8 10 8C8.89543 8 8 8.89543 8 10C8 11.1046 8.89543 12 10 12ZM10 11C9.44772 11 9 10.5523 9 10C9 9.44772 9.44772 9 10 9C10.5523 9 11 9.44772 11 10C11 10.5523 10.5523 11 10 11Z",fill:"currentColor"}))),Pe=n("rotateClockwise",i("svg",{viewBox:"0 0 20 20",fill:"none",xmlns:"http://www.w3.org/2000/svg"},i("path",{d:"M17 10C17 6.13401 13.866 3 10 3C6.13401 3 3 6.13401 3 10C3 12.7916 4.63419 15.2026 7 16.3265V14.5C7 14.2239 7.22386 14 7.5 14C7.77614 14 8 14.2239 8 14.5V17.5C8 17.7761 7.77614 18 7.5 18H4.5C4.22386 18 4 17.7761 4 17.5C4 17.2239 4.22386 17 4.5 17H6.12422C3.66539 15.6357 2 13.0128 2 10C2 5.58172 5.58172 2 10 2C14.4183 2 18 5.58172 18 10C18 10.2761 17.7761 10.5 17.5 10.5C17.2239 10.5 17 10.2761 17 10Z",fill:"currentColor"}),i("path",{d:"M10 12C8.89543 12 8 11.1046 8 10C8 8.89543 8.89543 8 10 8C11.1046 8 12 8.89543 12 10C12 11.1046 11.1046 12 10 12ZM10 11C10.5523 11 11 10.5523 11 10C11 9.44772 10.5523 9 10 9C9.44772 9 9 9.44772 9 10C9 10.5523 9.44772 11 10 11Z",fill:"currentColor"}))),Oe=n("zoomIn",i("svg",{viewBox:"0 0 20 20",fill:"none",xmlns:"http://www.w3.org/2000/svg"},i("path",{d:"M11.5 8.5C11.5 8.22386 11.2761 8 11 8H9V6C9 5.72386 8.77614 5.5 8.5 5.5C8.22386 5.5 8 5.72386 8 6V8H6C5.72386 8 5.5 8.22386 5.5 8.5C5.5 8.77614 5.72386 9 6 9H8V11C8 11.2761 8.22386 11.5 8.5 11.5C8.77614 11.5 9 11.2761 9 11V9H11C11.2761 9 11.5 8.77614 11.5 8.5Z",fill:"currentColor"}),i("path",{d:"M8.5 3C11.5376 3 14 5.46243 14 8.5C14 9.83879 13.5217 11.0659 12.7266 12.0196L16.8536 16.1464C17.0488 16.3417 17.0488 16.6583 16.8536 16.8536C16.68 17.0271 16.4106 17.0464 16.2157 16.9114L16.1464 16.8536L12.0196 12.7266C11.0659 13.5217 9.83879 14 8.5 14C5.46243 14 3 11.5376 3 8.5C3 5.46243 5.46243 3 8.5 3ZM8.5 4C6.01472 4 4 6.01472 4 8.5C4 10.9853 6.01472 13 8.5 13C10.9853 13 13 10.9853 13 8.5C13 6.01472 10.9853 4 8.5 4Z",fill:"currentColor"}))),Se=n("zoomOut",i("svg",{viewBox:"0 0 20 20",fill:"none",xmlns:"http://www.w3.org/2000/svg"},i("path",{d:"M11 8C11.2761 8 11.5 8.22386 11.5 8.5C11.5 8.77614 11.2761 9 11 9H6C5.72386 9 5.5 8.77614 5.5 8.5C5.5 8.22386 5.72386 8 6 8H11Z",fill:"currentColor"}),i("path",{d:"M14 8.5C14 5.46243 11.5376 3 8.5 3C5.46243 3 3 5.46243 3 8.5C3 11.5376 5.46243 14 8.5 14C9.83879 14 11.0659 13.5217 12.0196 12.7266L16.1464 16.8536L16.2157 16.9114C16.4106 17.0464 16.68 17.0271 16.8536 16.8536C17.0488 16.6583 17.0488 16.3417 16.8536 16.1464L12.7266 12.0196C13.5217 11.0659 14 9.83879 14 8.5ZM4 8.5C4 6.01472 6.01472 4 8.5 4C10.9853 4 13 6.01472 13 8.5C13 10.9853 10.9853 13 8.5 13C6.01472 13 4 10.9853 4 8.5Z",fill:"currentColor"}))),Le=r({name:"ResizeSmall",render:()=>i("svg",{xmlns:"http://www.w3.org/2000/svg",viewBox:"0 0 20 20"},i("g",{fill:"none"},i("path",{d:"M5.5 4A1.5 1.5 0 0 0 4 5.5v1a.5.5 0 0 1-1 0v-1A2.5 2.5 0 0 1 5.5 3h1a.5.5 0 0 1 0 1h-1zM16 5.5A1.5 1.5 0 0 0 14.5 4h-1a.5.5 0 0 1 0-1h1A2.5 2.5 0 0 1 17 5.5v1a.5.5 0 0 1-1 0v-1zm0 9a1.5 1.5 0 0 1-1.5 1.5h-1a.5.5 0 0 0 0 1h1a2.5 2.5 0 0 0 2.5-2.5v-1a.5.5 0 0 0-1 0v1zm-12 0A1.5 1.5 0 0 0 5.5 16h1.25a.5.5 0 0 1 0 1H5.5A2.5 2.5 0 0 1 3 14.5v-1.25a.5.5 0 0 1 1 0v1.25zM8.5 7A1.5 1.5 0 0 0 7 8.5v3A1.5 1.5 0 0 0 8.5 13h3a1.5 1.5 0 0 0 1.5-1.5v-3A1.5 1.5 0 0 0 11.5 7h-3zM8 8.5a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1-.5-.5v-3z",fill:"currentColor"})))}),ze=Object.assign(Object.assign({},l.props),{showToolbar:{type:Boolean,default:!0},showToolbarTooltip:Boolean}),Ie=a("n-image");const ke=s({name:"Image",common:u,peers:{Tooltip:c},self:function(){return{toolbarIconColor:"rgba(255, 255, 255, .9)",toolbarColor:"rgba(0, 0, 0, .35)",toolbarBoxShadow:"none",toolbarBorderRadius:"24px"}}}),Te=i("svg",{viewBox:"0 0 20 20",fill:"none",xmlns:"http://www.w3.org/2000/svg"},i("path",{d:"M6 5C5.75454 5 5.55039 5.17688 5.50806 5.41012L5.5 5.5V14.5C5.5 14.7761 5.72386 15 6 15C6.24546 15 6.44961 14.8231 6.49194 14.5899L6.5 14.5V5.5C6.5 5.22386 6.27614 5 6 5ZM13.8536 5.14645C13.68 4.97288 13.4106 4.9536 13.2157 5.08859L13.1464 5.14645L8.64645 9.64645C8.47288 9.82001 8.4536 10.0894 8.58859 10.2843L8.64645 10.3536L13.1464 14.8536C13.3417 15.0488 13.6583 15.0488 13.8536 14.8536C14.0271 14.68 14.0464 14.4106 13.9114 14.2157L13.8536 14.1464L9.70711 10L13.8536 5.85355C14.0488 5.65829 14.0488 5.34171 13.8536 5.14645Z",fill:"currentColor"})),ye=i("svg",{viewBox:"0 0 20 20",fill:"none",xmlns:"http://www.w3.org/2000/svg"},i("path",{d:"M13.5 5C13.7455 5 13.9496 5.17688 13.9919 5.41012L14 5.5V14.5C14 14.7761 13.7761 15 13.5 15C13.2545 15 13.0504 14.8231 13.0081 14.5899L13 14.5V5.5C13 5.22386 13.2239 5 13.5 5ZM5.64645 5.14645C5.82001 4.97288 6.08944 4.9536 6.28431 5.08859L6.35355 5.14645L10.8536 9.64645C11.0271 9.82001 11.0464 10.0894 10.9114 10.2843L10.8536 10.3536L6.35355 14.8536C6.15829 15.0488 5.84171 15.0488 5.64645 14.8536C5.47288 14.68 5.4536 14.4106 5.58859 14.2157L5.64645 14.1464L9.79289 10L5.64645 5.85355C5.45118 5.65829 5.45118 5.34171 5.64645 5.14645Z",fill:"currentColor"})),Me=i("svg",{viewBox:"0 0 20 20",fill:"none",xmlns:"http://www.w3.org/2000/svg"},i("path",{d:"M4.089 4.216l.057-.07a.5.5 0 0 1 .638-.057l.07.057L10 9.293l5.146-5.147a.5.5 0 0 1 .638-.057l.07.057a.5.5 0 0 1 .057.638l-.057.07L10.707 10l5.147 5.146a.5.5 0 0 1 .057.638l-.057.07a.5.5 0 0 1-.638.057l-.07-.057L10 10.707l-5.146 5.147a.5.5 0 0 1-.638.057l-.07-.057a.5.5 0 0 1-.057-.638l.057-.07L9.293 10L4.146 4.854a.5.5 0 0 1-.057-.638l.057-.07l-.057.07z",fill:"currentColor"})),Re=d([d("body >",[h("image-container","position: fixed;")]),h("image-preview-container","\n position: fixed;\n left: 0;\n right: 0;\n top: 0;\n bottom: 0;\n display: flex;\n "),h("image-preview-overlay","\n z-index: -1;\n position: absolute;\n left: 0;\n right: 0;\n top: 0;\n bottom: 0;\n background: rgba(0, 0, 0, .3);\n ",[f()]),h("image-preview-toolbar","\n z-index: 1;\n position: absolute;\n left: 50%;\n transform: translateX(-50%);\n border-radius: var(--n-toolbar-border-radius);\n height: 48px;\n bottom: 40px;\n padding: 0 12px;\n background: var(--n-toolbar-color);\n box-shadow: var(--n-toolbar-box-shadow);\n color: var(--n-toolbar-icon-color);\n transition: color .3s var(--n-bezier);\n display: flex;\n align-items: center;\n ",[h("base-icon","\n padding: 0 8px;\n font-size: 28px;\n cursor: pointer;\n "),f()]),h("image-preview-wrapper","\n position: absolute;\n left: 0;\n right: 0;\n top: 0;\n bottom: 0;\n display: flex;\n pointer-events: none;\n ",[v()]),h("image-preview","\n user-select: none;\n -webkit-user-select: none;\n pointer-events: all;\n margin: auto;\n max-height: calc(100vh - 32px);\n max-width: calc(100vw - 32px);\n transition: transform .3s var(--n-bezier);\n "),h("image","\n display: inline-flex;\n max-height: 100%;\n max-width: 100%;\n ",[p("preview-disabled","\n cursor: pointer;\n "),d("img","\n border-radius: inherit;\n ")])]),Ae=r({name:"ImagePreview",props:Object.assign(Object.assign({},ze),{onNext:Function,onPrev:Function,clsPrefix:{type:String,required:!0}}),setup(e){const t=l("Image","-image",Re,ke,e,g(e,"clsPrefix"));let o=null;const n=w(null),r=w(null),a=w(void 0),s=w(!1),u=w(!1),{localeRef:c}=N("Image");function d(t){var o,n;switch(t.key){case" ":t.preventDefault();break;case"ArrowLeft":null===(o=e.onPrev)||void 0===o||o.call(e);break;case"ArrowRight":null===(n=e.onNext)||void 0===n||n.call(e);break;case"Escape":X()}}m(s,(e=>{e?C("keydown",document,d):b("keydown",document,d)})),x((()=>{b("keydown",document,d)}));let h=0,f=0,v=0,p=0,I=0,k=0,T=0,y=0,M=!1;function R(e){const{clientX:t,clientY:o}=e;v=t-h,p=o-f,H(Y)}function A(e){const{value:t}=n;if(!t)return{offsetX:0,offsetY:0};const o=t.getBoundingClientRect(),{moveVerticalDirection:i,moveHorizontalDirection:r,deltaHorizontal:l,deltaVertical:a}=e||{};let s=0,u=0;return s=o.width<=window.innerWidth?0:o.left>0?(o.width-window.innerWidth)/2:o.right<window.innerWidth?-(o.width-window.innerWidth)/2:"horizontalRight"===r?Math.min((o.width-window.innerWidth)/2,I-(null!=l?l:0)):Math.max(-(o.width-window.innerWidth)/2,I-(null!=l?l:0)),u=o.height<=window.innerHeight?0:o.top>0?(o.height-window.innerHeight)/2:o.bottom<window.innerHeight?-(o.height-window.innerHeight)/2:"verticalBottom"===i?Math.min((o.height-window.innerHeight)/2,k-(null!=a?a:0)):Math.max(-(o.height-window.innerHeight)/2,k-(null!=a?a:0)),{offsetX:s,offsetY:u}}function Z(e){b("mousemove",document,R),b("mouseup",document,Z);const{clientX:t,clientY:o}=e;M=!1;const n=function(e){const{mouseUpClientX:t,mouseUpClientY:o,mouseDownClientX:n,mouseDownClientY:i}=e,r=n-t,l=i-o;return{moveVerticalDirection:"vertical"+(l>0?"Top":"Bottom"),moveHorizontalDirection:"horizontal"+(r>0?"Left":"Right"),deltaHorizontal:r,deltaVertical:l}}({mouseUpClientX:t,mouseUpClientY:o,mouseDownClientX:T,mouseDownClientY:y}),i=A(n);v=i.offsetX,p=i.offsetY,Y()}const j=P(Ie,null);let $=0,B=1,V=0;function W(){B=1,$=0}function U(){const{value:e}=n;if(!e)return 1;const{innerWidth:t,innerHeight:o}=window,i=e.naturalHeight/(o-32),r=e.naturalWidth/(t-32);return i<1&&r<1?1:Math.max(i,r)}function Y(e=!0){const{value:t}=n;if(!t)return;const{style:o}=t,i=E(null==j?void 0:j.previewedImgPropsRef.value.style);let r="";if("string"==typeof i)r=i+";";else for(const n in i)r+=`${be(n)}: ${i[n]};`;const l=`transform-origin: center; transform: translateX(${v}px) translateY(${p}px) rotate(${V}deg) scale(${B});`;o.cssText=M?r+"cursor: grabbing; transition: none;"+l:r+"cursor: grab;"+l+(e?"":"transition: none;"),e||t.offsetHeight}function X(){s.value=!s.value,u.value=!0}const F={setPreviewSrc:e=>{a.value=e},setThumbnailEl:e=>{o=e},toggleShow:X};const G=O((()=>{const{common:{cubicBezierEaseInOut:e},self:{toolbarIconColor:o,toolbarBorderRadius:n,toolbarBoxShadow:i,toolbarColor:r}}=t.value;return{"--n-bezier":e,"--n-toolbar-icon-color":o,"--n-toolbar-color":r,"--n-toolbar-border-radius":n,"--n-toolbar-box-shadow":i}})),{inlineThemeDisabled:q}=S(),J=q?L("image-preview",void 0,G,e):void 0;return Object.assign({previewRef:n,previewWrapperRef:r,previewSrc:a,show:s,appear:z(),displayed:u,previewedImgProps:null==j?void 0:j.previewedImgPropsRef,handleWheel(e){e.preventDefault()},handlePreviewMousedown:function(e){var t,o;if(null===(o=null==j?void 0:(t=j.previewedImgPropsRef.value).onMousedown)||void 0===o||o.call(t,e),0!==e.button)return;const{clientX:n,clientY:i}=e;M=!0,h=n-v,f=i-p,I=v,k=p,T=n,y=i,Y(),C("mousemove",document,R),C("mouseup",document,Z)},handlePreviewDblclick:function(e){var t,o;null===(o=null==j?void 0:(t=j.previewedImgPropsRef.value).onDblclick)||void 0===o||o.call(t,e);const n=U();B=B===n?1:n,Y()},syncTransformOrigin:function(){const{value:e}=r;if(!o||!e)return;const{style:t}=e,n=o.getBoundingClientRect(),i=n.left+n.width/2,l=n.top+n.height/2;t.transformOrigin=`${i}px ${l}px`},handleAfterLeave:()=>{W(),V=0,u.value=!1},handleDragStart:e=>{var t,o;null===(o=null==j?void 0:(t=j.previewedImgPropsRef.value).onDragstart)||void 0===o||o.call(t,e),e.preventDefault()},zoomIn:function(){const e=function(){const{value:e}=n;if(!e)return 1;const{innerWidth:t,innerHeight:o}=window,i=Math.max(1,e.naturalHeight/(o-32)),r=Math.max(1,e.naturalWidth/(t-32));return Math.max(3,2*i,2*r)}();B<e&&($+=1,B=Math.min(e,Math.pow(1.5,$)),Y())},zoomOut:function(){if(B>.5){const e=B;$-=1,B=Math.max(.5,Math.pow(1.5,$));const t=e-B;Y(!1);const o=A();B+=t,Y(!1),B-=t,v=o.offsetX,p=o.offsetY,Y()}},rotateCounterclockwise:function(){V-=90,Y()},rotateClockwise:function(){V+=90,Y()},handleSwitchPrev:function(){var t;W(),V=0,null===(t=e.onPrev)||void 0===t||t.call(e)},handleSwitchNext:function(){var t;W(),V=0,null===(t=e.onNext)||void 0===t||t.call(e)},withTooltip:function(o,n){if(e.showToolbarTooltip){const{value:e}=t;return i(D,{to:!1,theme:e.peers.Tooltip,themeOverrides:e.peerOverrides.Tooltip,keepAliveOnHover:!1},{default:()=>c.value[n],trigger:()=>o})}return o},resizeToOrignalImageSize:function(){B=U(),$=Math.ceil(Math.log(B)/Math.log(1.5)),v=0,p=0,Y()},cssVars:q?void 0:G,themeClass:null==J?void 0:J.themeClass,onRender:null==J?void 0:J.onRender},F)},render(){var e,t;const{clsPrefix:o}=this;return i(M,null,null===(t=(e=this.$slots).default)||void 0===t?void 0:t.call(e),i(I,{show:this.show},{default:()=>{var e;return this.show||this.displayed?(null===(e=this.onRender)||void 0===e||e.call(this),k(i("div",{class:[`${o}-image-preview-container`,this.themeClass],style:this.cssVars,onWheel:this.handleWheel},i(y,{name:"fade-in-transition",appear:this.appear},{default:()=>this.show?i("div",{class:`${o}-image-preview-overlay`,onClick:this.toggleShow}):null}),this.showToolbar?i(y,{name:"fade-in-transition",appear:this.appear},{default:()=>{if(!this.show)return null;const{withTooltip:e}=this;return i("div",{class:`${o}-image-preview-toolbar`},this.onPrev?i(M,null,e(i(R,{clsPrefix:o,onClick:this.handleSwitchPrev},{default:()=>Te}),"tipPrevious"),e(i(R,{clsPrefix:o,onClick:this.handleSwitchNext},{default:()=>ye}),"tipNext")):null,e(i(R,{clsPrefix:o,onClick:this.rotateCounterclockwise},{default:()=>i(Pe,null)}),"tipCounterclockwise"),e(i(R,{clsPrefix:o,onClick:this.rotateClockwise},{default:()=>i(xe,null)}),"tipClockwise"),e(i(R,{clsPrefix:o,onClick:this.resizeToOrignalImageSize},{default:()=>i(Le,null)}),"tipOriginalSize"),e(i(R,{clsPrefix:o,onClick:this.zoomOut},{default:()=>i(Se,null)}),"tipZoomOut"),e(i(R,{clsPrefix:o,onClick:this.zoomIn},{default:()=>i(Oe,null)}),"tipZoomIn"),e(i(R,{clsPrefix:o,onClick:this.toggleShow},{default:()=>Me}),"tipClose"))}}):null,i(y,{name:"fade-in-scale-up-transition",onAfterLeave:this.handleAfterLeave,appear:this.appear,onEnter:this.syncTransformOrigin,onBeforeLeave:this.syncTransformOrigin},{default:()=>{const{previewedImgProps:e={}}=this;return k(i("div",{class:`${o}-image-preview-wrapper`,ref:"previewWrapperRef"},i("img",Object.assign({},e,{draggable:!1,onMousedown:this.handlePreviewMousedown,onDblclick:this.handlePreviewDblclick,class:[`${o}-image-preview`,e.class],key:this.previewSrc,src:this.previewSrc,ref:"previewRef",onDragstart:this.handleDragStart}))),[[A,this.show]])}})),[[T,{enabled:this.show}]])):null}}))}}),Ee=a("n-image-group"),De=r({name:"ImageGroup",props:ze,setup(e){let t;const{mergedClsPrefixRef:o}=S(e),n=`c${Z()}`,i=$(),r=e=>{var o;t=e,null===(o=a.value)||void 0===o||o.setPreviewSrc(e)};function l(e){if(!(null==i?void 0:i.proxy))return;const o=i.proxy.$el.parentElement.querySelectorAll(`[data-group-id=${n}]:not([data-error=true])`);if(!o.length)return;const l=Array.from(o).findIndex((e=>e.dataset.previewSrc===t));r(~l?o[(l+e+o.length)%o.length].dataset.previewSrc:o[0].dataset.previewSrc)}j(Ee,{mergedClsPrefixRef:o,setPreviewSrc:r,setThumbnailEl:e=>{var t;null===(t=a.value)||void 0===t||t.setThumbnailEl(e)},toggleShow:()=>{var e;null===(e=a.value)||void 0===e||e.toggleShow()},groupId:n});const a=w(null);return{mergedClsPrefix:o,previewInstRef:a,next:()=>l(1),prev:()=>l(-1)}},render(){return i(Ae,{theme:this.theme,themeOverrides:this.themeOverrides,clsPrefix:this.mergedClsPrefix,ref:"previewInstRef",onPrev:this.prev,onNext:this.next,showToolbar:this.showToolbar,showToolbarTooltip:this.showToolbarTooltip},this.$slots)}}),He=r({name:"Image",props:Object.assign({alt:String,height:[String,Number],imgProps:Object,previewedImgProps:Object,lazy:Boolean,intersectionObserverOptions:Object,objectFit:{type:String,default:"fill"},previewSrc:String,fallbackSrc:String,width:[String,Number],src:String,previewDisabled:Boolean,loadDescription:String,onError:Function,onLoad:Function},ze),inheritAttrs:!1,setup(o){const n=w(null),i=w(!1),r=w(null),l=P(Ee,null),{mergedClsPrefixRef:a}=l||S(o),s={click:()=>{if(o.previewDisabled||i.value)return;const e=o.previewSrc||o.src;if(l)return l.setPreviewSrc(e),l.setThumbnailEl(n.value),void l.toggleShow();const{value:t}=r;t&&(t.setPreviewSrc(e),t.setThumbnailEl(n.value),t.toggleShow())}},u=w(!o.lazy);B((()=>{var e;null===(e=n.value)||void 0===e||e.setAttribute("data-group-id",(null==l?void 0:l.groupId)||"")})),B((()=>{if(e)return;let i;const r=V((()=>{null==i||i(),i=void 0,o.lazy&&(i=t(n.value,o.intersectionObserverOptions,u))}));x((()=>{r(),null==i||i()}))})),V((()=>{var e;o.src,null===(e=o.imgProps)||void 0===e||e.src,i.value=!1}));const c=w(!1);return j(Ie,{previewedImgPropsRef:g(o,"previewedImgProps")}),Object.assign({mergedClsPrefix:a,groupId:null==l?void 0:l.groupId,previewInstRef:r,imageRef:n,showError:i,shouldStartLoading:u,loaded:c,mergedOnClick:e=>{var t,n;s.click(),null===(n=null===(t=o.imgProps)||void 0===t?void 0:t.onClick)||void 0===n||n.call(t,e)},mergedOnError:e=>{if(!u.value)return;i.value=!0;const{onError:t,imgProps:{onError:n}={}}=o;null==t||t(e),null==n||n(e)},mergedOnLoad:e=>{const{onLoad:t,imgProps:{onLoad:n}={}}=o;null==t||t(e),null==n||n(e),c.value=!0}},s)},render(){var t,o;const{mergedClsPrefix:n,imgProps:r={},loaded:l,$attrs:a,lazy:s}=this,u=null===(o=(t=this.$slots).placeholder)||void 0===o?void 0:o.call(t),c=this.src||r.src||"",d=i("img",Object.assign(Object.assign({},r),{ref:"imageRef",width:this.width||r.width,height:this.height||r.height,src:e?c:this.showError?this.fallbackSrc:this.shouldStartLoading?c:void 0,alt:this.alt||r.alt,"aria-label":this.alt||r.alt,onClick:this.mergedOnClick,onError:this.mergedOnError,onLoad:this.mergedOnLoad,loading:s?"lazy":"eager",style:[r.style||"",u&&!l?{height:"0",width:"0",visibility:"hidden"}:"",{objectFit:this.objectFit}],"data-error":this.showError,"data-preview-src":this.previewSrc||this.src}));return i("div",Object.assign({},a,{role:"none",class:[a.class,`${n}-image`,(this.previewDisabled||this.showError)&&`${n}-image--preview-disabled`]}),this.groupId?d:i(Ae,{theme:this.theme,themeOverrides:this.themeOverrides,clsPrefix:n,ref:"previewInstRef",showToolbar:this.showToolbar,showToolbarTooltip:this.showToolbarTooltip},{default:()=>d}),!l&&u)}});export{He as N,De as a};