var G=Object.defineProperty;var B=(i,t,n)=>t in i?G(i,t,{enumerable:!0,configurable:!0,writable:!0,value:n}):i[t]=n;var p=(i,t,n)=>B(i,typeof t!="symbol"?t+"":t,n);(function(){const t=document.createElement("link").relList;if(t&&t.supports&&t.supports("modulepreload"))return;for(const s of document.querySelectorAll('link[rel="modulepreload"]'))e(s);new MutationObserver(s=>{for(const r of s)if(r.type==="childList")for(const o of r.addedNodes)o.tagName==="LINK"&&o.rel==="modulepreload"&&e(o)}).observe(document,{childList:!0,subtree:!0});function n(s){const r={};return s.integrity&&(r.integrity=s.integrity),s.referrerPolicy&&(r.referrerPolicy=s.referrerPolicy),s.crossOrigin==="use-credentials"?r.credentials="include":s.crossOrigin==="anonymous"?r.credentials="omit":r.credentials="same-origin",r}function e(s){if(s.ep)return;s.ep=!0;const r=n(s);fetch(s.href,r)}})();class P{constructor(){this._events={}}on(t,n,e={}){this._events[t]=this._events[t]||[],this._events[t].push({fn:n,options:e})}off(t,n){const e=this._events[t]||[];this._events[t]=e.filter(s=>s.fn!==n)}find(t){return this._events[t]}run(t,...n){const e=this.getSubscribers(t,this._events);return console.assert(e&&e.length>0,"No subscriber for event: "+t),e.forEach(s=>{const{fn:r,options:o}=s;return o.delay?this.delay(t,r,n,o):Object.keys(o).length>0?r.apply(this,[...n,o]):r.apply(this,n),!s.options.once}),e.length}once(t,n,e={}){this.on(t,n,Object.assign(Object.assign({},e),{once:!0}))}delay(t,n,e,s){s._t&&clearTimeout(s._t),s._t=setTimeout(()=>{clearTimeout(s._t),Object.keys(s).length>0?n.apply(this,[...e,s]):n.apply(this,e)},s.delay)}runAsync(t,...n){const e=this.getSubscribers(t,this._events);console.assert(e&&e.length>0,"No subscriber for event: "+t);const s=e.map(r=>{const{fn:o,options:c}=r;return Object.keys(c).length>0?o.apply(this,[...n,c]):o.apply(this,n)});return Promise.all(s)}query(t,...n){return this.runAsync(t,...n)}getSubscribers(t,n){const e=n[t]||[];return n[t]=e.filter(s=>!s.options.once),Object.keys(n).filter(s=>s.endsWith("*")&&t.startsWith(s.replace("*",""))).sort((s,r)=>r.length-s.length).forEach(s=>e.push(...n[s].map(r=>Object.assign(Object.assign({},r),{options:Object.assign(Object.assign({},r.options),{event:t})})))),e}}const U="AppRun-3";let m;const b=typeof self=="object"&&self.self===self&&self||typeof global=="object"&&global.global===global&&global;b.app&&b._AppRunVersions?m=b.app:(m=new P,b.app=m,b._AppRunVersions=U);const a=m,d=(i,t)=>(t?i.state[t]:i.state)||"",_=(i,t,n)=>{if(t){const e=i.state||{};e[t]=n,i.setState(e)}else i.setState(n)},I=(i,t,n,e)=>{if(i.startsWith("$on")){const s=t[i];if(i=i.substring(1),typeof s=="boolean")t[i]=r=>e.run?e.run(i,r):a.run(i,r);else if(typeof s=="string")t[i]=r=>e.run?e.run(s,r):a.run(s,r);else if(typeof s=="function")t[i]=r=>e.setState(s(e.state,r));else if(Array.isArray(s)){const[r,...o]=s;typeof r=="string"?t[i]=c=>e.run?e.run(r,...o,c):a.run(r,...o,c):typeof r=="function"&&(t[i]=c=>e.setState(r(e.state,...o,c)))}}else if(i==="$bind"){const s=t.type||"text",r=typeof t[i]=="string"?t[i]:t.name;if(n==="input")switch(s){case"checkbox":t.checked=d(e,r),t.onclick=o=>_(e,r||o.target.name,o.target.checked);break;case"radio":t.checked=d(e,r)===t.value,t.onclick=o=>_(e,r||o.target.name,o.target.value);break;case"number":case"range":t.value=d(e,r),t.oninput=o=>_(e,r||o.target.name,Number(o.target.value));break;default:t.value=d(e,r),t.oninput=o=>_(e,r||o.target.name,o.target.value)}else n==="select"?(t.value=d(e,r),t.onchange=o=>{o.target.multiple||_(e,r||o.target.name,o.target.value)}):n==="option"?(t.selected=d(e,r),t.onclick=o=>_(e,r||o.target.name,o.target.selected)):n==="textarea"&&(t.innerHTML=d(e,r),t.oninput=o=>_(e,r||o.target.name,o.target.value))}else a.run("$",{key:i,tag:n,props:t,component:e})},A=(i,t)=>{if(Array.isArray(i))return i.map(n=>A(n,t));{let{type:n,tag:e,props:s,children:r}=i;return e=e||n,r=r||(s==null?void 0:s.children),s&&Object.keys(s).forEach(o=>{o.startsWith("$")&&(I(o,s,e,t),delete s[o])}),r&&A(r,t),i}};function D(i,...t){return F(t)}const V="_props";function F(i){const t=[],n=e=>{e!=null&&e!==""&&e!==!1&&t.push(typeof e=="function"||typeof e=="object"?e:`${e}`)};return i&&i.forEach(e=>{Array.isArray(e)?e.forEach(s=>n(s)):n(e)}),t}function K(i,t,...n){const e=F(n);if(typeof i=="string")return{tag:i,props:t,children:e};if(Array.isArray(i))return i;if(i===void 0&&n)return e;if(Object.getPrototypeOf(i).__isAppRunComponent)return{tag:i,props:t,children:e};if(typeof i=="function")return i(t,e);throw new Error(`Unknown tag in vdom ${i}`)}const H=new WeakMap,J=(i,t,n={})=>{if(t==null||t===!1)return;const e=typeof i=="string"&&i?document.getElementById(i)||document.querySelector(i):i;t=A(t,n),Q(e,t,n)};function Q(i,t,n={}){if(t==null||t===!1||(t=g(t,n),!i))return;const e=i.nodeName==="SVG";Array.isArray(t)?k(i,t,e):k(i,[t],e)}function z(i,t){const n=i.nodeName,e=`${t.tag||""}`;return n.toUpperCase()===e.toUpperCase()}function S(i,t,n){if(t._op!==3){if(n=n||t.tag==="svg",!z(i,t)){i.parentNode.replaceChild(E(t,n),i);return}!(t._op&2)&&k(i,t.children,n),!(t._op&1)&&$(i,t.props,n)}}function k(i,t,n){var e,s;const r=((e=i.childNodes)===null||e===void 0?void 0:e.length)||0,o=(t==null?void 0:t.length)||0,c=Math.min(r,o);for(let u=0;u<c;u++){const h=t[u];if(h._op===3)continue;const f=i.childNodes[u];if(typeof h=="string")f.textContent!==h&&(f.nodeType===3?f.nodeValue=h:i.replaceChild(N(h),f));else if(h instanceof HTMLElement||h instanceof SVGElement)i.insertBefore(h,f);else{const C=h.props&&h.props.key;if(C)if(f.key===C)S(i.childNodes[u],h,n);else{const j=H[C];if(j){const L=j.nextSibling;i.insertBefore(j,f),L?i.insertBefore(f,L):i.appendChild(f),S(i.childNodes[u],h,n)}else i.replaceChild(E(h,n),f)}else S(i.childNodes[u],h,n)}}let l=((s=i.childNodes)===null||s===void 0?void 0:s.length)||0;for(;l>c;)i.removeChild(i.lastChild),l--;if(o>c){const u=document.createDocumentFragment();for(let h=c;h<t.length;h++)u.appendChild(E(t[h],n));i.appendChild(u)}}const R=i=>{const t=document.createElement("section");return t.insertAdjacentHTML("afterbegin",i),Array.from(t.children)};function N(i){if((i==null?void 0:i.indexOf("_html:"))===0){const t=document.createElement("div");return t.insertAdjacentHTML("afterbegin",i.substring(6)),t}else return document.createTextNode(i??"")}function E(i,t){if(i instanceof HTMLElement||i instanceof SVGElement)return i;if(typeof i=="string")return N(i);if(!i.tag||typeof i.tag=="function")return N(JSON.stringify(i));t=t||i.tag==="svg";const n=t?document.createElementNS("http://www.w3.org/2000/svg",i.tag):document.createElement(i.tag);return $(n,i.props,t),i.children&&i.children.forEach(e=>n.appendChild(E(e,t))),n}function X(i,t){t.class=t.class||t.className,delete t.className;const n={};return i&&Object.keys(i).forEach(e=>n[e]=null),Object.keys(t).forEach(e=>n[e]=t[e]),n}function $(i,t,n){const e=i[V]||{};t=X(e,t||{}),i[V]=t;for(const s in t){const r=t[s];if(s.startsWith("data-")){const c=s.substring(5).replace(/-(\w)/g,l=>l[1].toUpperCase());i.dataset[c]!==r&&(r||r===""?i.dataset[c]=r:delete i.dataset[c])}else if(s==="style")if(i.style.cssText&&(i.style.cssText=""),typeof r=="string")i.style.cssText=r;else for(const o in r)i.style[o]!==r[o]&&(i.style[o]=r[o]);else if(s.startsWith("xlink")){const o=s.replace("xlink","").toLowerCase();r==null||r===!1?i.removeAttributeNS("http://www.w3.org/1999/xlink",o):i.setAttributeNS("http://www.w3.org/1999/xlink",o,r)}else s.startsWith("on")?!r||typeof r=="function"?i[s]=r:typeof r=="string"&&(r?i.setAttribute(s,r):i.removeAttribute(s)):/^id$|^class$|^list$|^readonly$|^contenteditable$|^role|-|^for$/g.test(s)||n?i.getAttribute(s)!==r&&(r?i.setAttribute(s,r):i.removeAttribute(s)):i[s]!==r&&(i[s]=r);s==="key"&&r&&(H[r]=i)}t&&typeof t.ref=="function"&&window.requestAnimationFrame(()=>t.ref(i))}function Y(i,t,n){const{tag:e,props:s,children:r}=i;let o=`_${n}`,c=s&&s.id;c?o=c:c=`_${n}${Date.now()}`;let l="section";s&&s.as&&(l=s.as,delete s.as),t.__componentCache||(t.__componentCache={});let u=t.__componentCache[o];if(!u||!(u instanceof e)||!u.element){const h=document.createElement(l);u=t.__componentCache[o]=new e(Object.assign(Object.assign({},s),{children:r})).mount(h,{render:!0})}else u.renderState(u.state);if(u.mounted){const h=u.mounted(s,r,u.state);typeof h<"u"&&u.setState(h)}return $(u.element,s,!1),u.element}function g(i,t,n=0){var e;if(typeof i=="string")return i;if(Array.isArray(i))return i.map(r=>g(r,t,n++));let s=i;if(i&&typeof i.tag=="function"&&Object.getPrototypeOf(i.tag).__isAppRunComponent&&(s=Y(i,t,n)),s&&Array.isArray(s.children)){const r=(e=s.props)===null||e===void 0?void 0:e._component;if(r){let o=0;s.children=s.children.map(c=>g(c,r,o++))}else s.children=s.children.map(o=>g(o,t,n++))}return s}const Z=(i,t={})=>class extends HTMLElement{constructor(){super()}get component(){return this._component}get state(){return this._component.state}static get observedAttributes(){return(t.observedAttributes||[]).map(e=>e.toLowerCase())}connectedCallback(){if(this.isConnected&&!this._component){const e=t||{};this._shadowRoot=e.shadow?this.attachShadow({mode:"open"}):this;const s=e.observedAttributes||[],r=s.reduce((c,l)=>{const u=l.toLowerCase();return u!==l&&(c[u]=l),c},{});this._attrMap=c=>r[c]||c;const o={};Array.from(this.attributes).forEach(c=>o[this._attrMap(c.name)]=c.value),s.forEach(c=>{this[c]!==void 0&&(o[c]=this[c]),Object.defineProperty(this,c,{get(){return o[c]},set(l){this.attributeChangedCallback(c,o[c],l)},configurable:!0,enumerable:!0})}),requestAnimationFrame(()=>{const c=this.children?Array.from(this.children):[];if(this._component=new i(Object.assign(Object.assign({},o),{children:c})).mount(this._shadowRoot,e),this._component._props=o,this._component.dispatchEvent=this.dispatchEvent.bind(this),this._component.mounted){const l=this._component.mounted(o,c,this._component.state);typeof l<"u"&&(this._component.state=l)}this.on=this._component.on.bind(this._component),this.run=this._component.run.bind(this._component),e.render!==!1&&this._component.run(".")})}}disconnectedCallback(){var e,s,r,o;(s=(e=this._component)===null||e===void 0?void 0:e.unload)===null||s===void 0||s.call(e),(o=(r=this._component)===null||r===void 0?void 0:r.unmount)===null||o===void 0||o.call(r),this._component=null}attributeChangedCallback(e,s,r){if(this._component){const o=this._attrMap(e);this._component._props[o]=r,this._component.run("attributeChanged",o,s,r),r!==s&&t.render!==!1&&window.requestAnimationFrame(()=>{this._component.run(".")})}}},q=(i,t,n)=>{typeof customElements<"u"&&customElements.define(i,Z(t,n))},M={meta:new WeakMap,defineMetadata(i,t,n){this.meta.has(n)||this.meta.set(n,{}),this.meta.get(n)[i]=t},getMetadataKeys(i){return i=Object.getPrototypeOf(i),this.meta.get(i)?Object.keys(this.meta.get(i)):[]},getMetadata(i,t){return t=Object.getPrototypeOf(t),this.meta.get(t)?this.meta.get(t)[i]:null}};function tt(i,t={}){return function(n,e){const s=i?i.toString():e;M.defineMetadata(`apprun-update:${s}`,{name:s,key:e,options:t},n)}}function et(i,t){return function(e){return q(i,e,t),e}}const v=new Map;a.find("get-components")||a.on("get-components",i=>i.components=v);const W=i=>i;class O{renderState(t,n=null){if(!this.view)return;let e=n||this.view(t);if(a.debug&&a.run("debug",{component:this,_:e?".":"-",state:t,vdom:e,el:this.element}),typeof document!="object")return;const s=typeof this.element=="string"&&this.element?document.getElementById(this.element)||document.querySelector(this.element):this.element;if(!s)return;const r="_c";this.unload?(s._component!==this||s.getAttribute(r)!==this.tracking_id)&&(this.tracking_id=new Date().valueOf().toString(),s.setAttribute(r,this.tracking_id),typeof MutationObserver<"u"&&(this.observer||(this.observer=new MutationObserver(o=>{(o[0].oldValue===this.tracking_id||!document.body.contains(s))&&(this.unload(this.state),this.observer.disconnect(),this.observer=null)})),this.observer.observe(document.body,{childList:!0,subtree:!0,attributes:!0,attributeOldValue:!0,attributeFilter:[r]}))):s.removeAttribute&&s.removeAttribute(r),s._component=this,!n&&e&&(e=A(e,this),this.options.transition&&document&&document.startViewTransition?document.startViewTransition(()=>a.render(s,e,this)):a.render(s,e,this)),this.rendered&&this.rendered(this.state)}setState(t,n={render:!0,history:!1}){if(t instanceof Promise)Promise.resolve(t).then(e=>{this.setState(e,n),this._state=t});else{if(this._state=t,t==null)return;this.state=t,n.render!==!1&&(n.transition&&document&&document.startViewTransition?document.startViewTransition(()=>this.renderState(t)):this.renderState(t)),n.history!==!1&&this.enable_history&&(this._history=[...this._history,t],this._history_idx=this._history.length-1),typeof n.callback=="function"&&n.callback(this.state)}}constructor(t,n,e,s){this.state=t,this.view=n,this.update=e,this.options=s,this._app=new P,this._actions=[],this._global_events=[],this._history=[],this._history_idx=-1,this._history_prev=()=>{this._history_idx--,this._history_idx>=0?this.setState(this._history[this._history_idx],{render:!0,history:!1}):this._history_idx=0},this._history_next=()=>{this._history_idx++,this._history_idx<this._history.length?this.setState(this._history[this._history_idx],{render:!0,history:!1}):this._history_idx=this._history.length-1},this.start=(r=null,o)=>{if(this.mount(r,Object.assign({render:!0},o)),this.mounted&&typeof this.mounted=="function"){const c=this.mounted({},[],this.state);typeof c<"u"&&this.setState(c)}return this}}mount(t=null,n){var e,s;return console.assert(!this.element,"Component already mounted."),this.options=n=Object.assign(Object.assign({},this.options),n),this.element=t,this.global_event=n.global_event,this.enable_history=!!n.history,this.enable_history&&(this.on(n.history.prev||"history-prev",this._history_prev),this.on(n.history.next||"history-next",this._history_next)),n.route&&(this.update=this.update||{},this.update[n.route]||(this.update[n.route]=W)),this.add_actions(),this.state=(s=(e=this.state)!==null&&e!==void 0?e:this.model)!==null&&s!==void 0?s:{},typeof this.state=="function"&&(this.state=this.state()),this.setState(this.state,{render:!!n.render,history:!0}),a.debug&&(v.get(t)?v.get(t).push(this):v.set(t,[this])),this}is_global_event(t){return t&&(this.global_event||this._global_events.indexOf(t)>=0||t.startsWith("#")||t.startsWith("/")||t.startsWith("@"))}add_action(t,n,e={}){!n||typeof n!="function"||(e.global&&this._global_events.push(t),this.on(t,(...s)=>{a.debug&&a.run("debug",{component:this,_:">",event:t,p:s,current_state:this.state,options:e});const r=n(this.state,...s);a.debug&&a.run("debug",{component:this,_:"<",event:t,p:s,newState:r,state:this.state,options:e}),this.setState(r,e)},e))}add_actions(){const t=this.update||{};M.getMetadataKeys(this).forEach(e=>{if(e.startsWith("apprun-update:")){const s=M.getMetadata(e,this);t[s.name]=[this[s.key].bind(this),s.options]}});const n={};Array.isArray(t)?t.forEach(e=>{const[s,r,o]=e;s.toString().split(",").forEach(l=>n[l.trim()]=[r,o])}):Object.keys(t).forEach(e=>{const s=t[e];(typeof s=="function"||Array.isArray(s))&&e.split(",").forEach(r=>n[r.trim()]=s)}),n["."]||(n["."]=W),Object.keys(n).forEach(e=>{const s=n[e];typeof s=="function"?this.add_action(e,s):Array.isArray(s)&&this.add_action(e,s[0],s[1])})}run(t,...n){if(this.state instanceof Promise)return Promise.resolve(this.state).then(e=>{this.state=e,this.run(t,...n)});{const e=t.toString();return this.is_global_event(e)?a.run(e,...n):this._app.run(e,...n)}}on(t,n,e){const s=t.toString();return this._actions.push({name:s,fn:n}),this.is_global_event(s)?a.on(s,n,e):this._app.on(s,n,e)}runAsync(t,...n){const e=t.toString();return this.is_global_event(e)?a.runAsync(e,...n):this._app.runAsync(e,...n)}query(t,...n){return this.runAsync(t,...n)}unmount(){var t;(t=this.observer)===null||t===void 0||t.disconnect(),this._actions.forEach(n=>{const{name:e,fn:s}=n;this.is_global_event(e)?a.off(e,s):this._app.off(e,s)})}}O.__isAppRunComponent=!0;const w="//",x="///",y=i=>{if(i||(i="#"),i.startsWith("#")){const[t,...n]=i.split("/");a.run(t,...n)||a.run(x,t,...n),a.run(w,t,...n)}else if(i.startsWith("/")){const[t,n,...e]=i.split("/");a.run("/"+n,...e)||a.run(x,"/"+n,...e),a.run(w,"/"+n,...e)}else a.run(i)||a.run(x,i),a.run(w,i)};if(!a.start){a.h=a.createElement=K,a.render=J,a.Fragment=D,a.webComponent=q,a.safeHTML=R,a.start=(t,n,e,s,r)=>{const o=Object.assign({render:!0,global_event:!0},r),c=new O(n,e,s);return r&&r.rendered&&(c.rendered=r.rendered),r&&r.mounted&&(c.mounted=r.mounted),c.start(t,o),c};const i=t=>{};a.on("$",i),a.on("debug",t=>i),a.on(w,i),a.on("#",i),a.route=y,a.on("route",t=>a.route&&a.route(t)),typeof document=="object"&&document.addEventListener("DOMContentLoaded",()=>{a.route===y&&(window.onpopstate=()=>y(location.hash),document.body.hasAttribute("apprun-no-init")||a["no-init-route"]||y(location.hash))}),typeof window=="object"&&(window.Component=O,window._React=window.React,window.React=a,window.on=tt,window.customElement=et,window.safeHTML=R),a.use_render=(t,n=0)=>n===0?a.render=(e,s)=>t(s,e):a.render=(e,s)=>t(e,s),a.use_react=(t,n)=>{a.h=a.createElement=t.createElement,a.Fragment=t.Fragment,a.render=(e,s)=>n.render(s,e),t.version&&t.version.startsWith("18")&&(a.render=(e,s)=>{!e||!s||(e._root||(e._root=n.createRoot(e)),e._root.render(s))})}}class it extends O{constructor(){super(...arguments);p(this,"state",{player:{name:"Adventurer",health:100,level:1,gold:0},currentView:"main"});p(this,"view",n=>`
      <div class="game-container">
        <header>
          <h1>GalyCherryGame</h1>
          <div class="player-stats">
            <div>Name: ${n.player.name}</div>
            <div>Health: ${n.player.health}</div>
            <div>Level: ${n.player.level}</div>
            <div>Gold: ${n.player.gold}</div>
          </div>
        </header>
        <main>
          ${this.getCurrentView(n)}
        </main>
      </div>
    `);p(this,"getCurrentView",n=>{switch(n.currentView){case"main":return`
          <div class="main-menu">
            <button onclick="app.run('navigate', 'combat')">Combat</button>
            <button onclick="app.run('navigate', 'crafting')">Crafting</button>
            <button onclick="app.run('navigate', 'quests')">Quests</button>
          </div>
        `;default:return"<div>Coming Soon</div>"}});p(this,"update",{navigate:(n,e)=>({...n,currentView:e})})}}const T=document.getElementById("app");if(T){const i=new it;i.mount(T),a.start(T,i.state,i.view,i.update)}