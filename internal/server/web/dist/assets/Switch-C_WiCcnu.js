import{An as e,Cr as t,Fn as n,Hn as r,Ir as i,Mn as a,Sr as o,Tr as s,Ur as c,Wr as l,Xt as u,Zt as d,_r as f,br as p,cn as m,dn as h,fr as g,hr as _,ii as v,jn as y,kn as b,ri as x,sr as S,tn as C,wr as w,xr as T}from"./index-Dmrq63M5.js";var E={buttonHeightSmall:`14px`,buttonHeightMedium:`18px`,buttonHeightLarge:`22px`,buttonWidthSmall:`14px`,buttonWidthMedium:`18px`,buttonWidthLarge:`22px`,buttonWidthPressedSmall:`20px`,buttonWidthPressedMedium:`24px`,buttonWidthPressedLarge:`28px`,railHeightSmall:`18px`,railHeightMedium:`22px`,railHeightLarge:`26px`,railWidthSmall:`32px`,railWidthMedium:`40px`,railWidthLarge:`48px`};function D(e){let{primaryColor:t,opacityDisabled:n,borderRadius:r,textColor3:i}=e;return Object.assign(Object.assign({},E),{iconColor:i,textColor:`white`,loadingColor:t,opacityDisabled:n,railColor:`rgba(0, 0, 0, .14)`,railColorActive:t,buttonBoxShadow:`0 1px 4px 0 rgba(0, 0, 0, 0.3), inset 0 0 1px 0 rgba(0, 0, 0, 0.05)`,buttonColor:`#FFF`,railBorderRadiusSmall:r,railBorderRadiusMedium:r,railBorderRadiusLarge:r,buttonBorderRadiusSmall:r,buttonBorderRadiusMedium:r,buttonBorderRadiusLarge:r,boxShadowFocus:`0 0 0 2px ${g(t,{alpha:.2})}`})}var O={name:`Switch`,common:u,self:D},k=T(`switch`,`
 height: var(--n-height);
 min-width: var(--n-width);
 vertical-align: middle;
 user-select: none;
 -webkit-user-select: none;
 display: inline-flex;
 outline: none;
 justify-content: center;
 align-items: center;
`,[o(`children-placeholder`,`
 height: var(--n-rail-height);
 display: flex;
 flex-direction: column;
 overflow: hidden;
 pointer-events: none;
 visibility: hidden;
 `),o(`rail-placeholder`,`
 display: flex;
 flex-wrap: none;
 `),o(`button-placeholder`,`
 width: calc(1.75 * var(--n-rail-height));
 height: var(--n-rail-height);
 `),T(`base-loading`,`
 position: absolute;
 top: 50%;
 left: 50%;
 transform: translateX(-50%) translateY(-50%);
 font-size: calc(var(--n-button-width) - 4px);
 color: var(--n-loading-color);
 transition: color .3s var(--n-bezier);
 `,[C({left:`50%`,top:`50%`,originalTransform:`translateX(-50%) translateY(-50%)`})]),o(`checked, unchecked`,`
 transition: color .3s var(--n-bezier);
 color: var(--n-text-color);
 box-sizing: border-box;
 position: absolute;
 white-space: nowrap;
 top: 0;
 bottom: 0;
 display: flex;
 align-items: center;
 line-height: 1;
 `),o(`checked`,`
 right: 0;
 padding-right: calc(1.25 * var(--n-rail-height) - var(--n-offset));
 `),o(`unchecked`,`
 left: 0;
 justify-content: flex-end;
 padding-left: calc(1.25 * var(--n-rail-height) - var(--n-offset));
 `),p(`&:focus`,[o(`rail`,`
 box-shadow: var(--n-box-shadow-focus);
 `)]),t(`round`,[o(`rail`,`border-radius: calc(var(--n-rail-height) / 2);`,[o(`button`,`border-radius: calc(var(--n-button-height) / 2);`)])]),w(`disabled`,[w(`icon`,[t(`rubber-band`,[t(`pressed`,[o(`rail`,[o(`button`,`max-width: var(--n-button-width-pressed);`)])]),o(`rail`,[p(`&:active`,[o(`button`,`max-width: var(--n-button-width-pressed);`)])]),t(`active`,[t(`pressed`,[o(`rail`,[o(`button`,`left: calc(100% - var(--n-offset) - var(--n-button-width-pressed));`)])]),o(`rail`,[p(`&:active`,[o(`button`,`left: calc(100% - var(--n-offset) - var(--n-button-width-pressed));`)])])])])])]),t(`active`,[o(`rail`,[o(`button`,`left: calc(100% - var(--n-button-width) - var(--n-offset))`)])]),o(`rail`,`
 overflow: hidden;
 height: var(--n-rail-height);
 min-width: var(--n-rail-width);
 border-radius: var(--n-rail-border-radius);
 cursor: pointer;
 position: relative;
 transition:
 opacity .3s var(--n-bezier),
 background .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier);
 background-color: var(--n-rail-color);
 `,[o(`button-icon`,`
 color: var(--n-icon-color);
 transition: color .3s var(--n-bezier);
 font-size: calc(var(--n-button-height) - 4px);
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 display: flex;
 justify-content: center;
 align-items: center;
 line-height: 1;
 `,[C()]),o(`button`,`
 align-items: center; 
 top: var(--n-offset);
 left: var(--n-offset);
 height: var(--n-button-height);
 width: var(--n-button-width-pressed);
 max-width: var(--n-button-width);
 border-radius: var(--n-button-border-radius);
 background-color: var(--n-button-color);
 box-shadow: var(--n-button-box-shadow);
 box-sizing: border-box;
 cursor: inherit;
 content: "";
 position: absolute;
 transition:
 background-color .3s var(--n-bezier),
 left .3s var(--n-bezier),
 opacity .3s var(--n-bezier),
 max-width .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier);
 `)]),t(`active`,[o(`rail`,`background-color: var(--n-rail-color-active);`)]),t(`loading`,[o(`rail`,`
 cursor: wait;
 `)]),t(`disabled`,[o(`rail`,`
 cursor: not-allowed;
 opacity: .5;
 `)])]),A=Object.assign(Object.assign({},h.props),{size:String,value:{type:[String,Number,Boolean],default:void 0},loading:Boolean,defaultValue:{type:[String,Number,Boolean],default:!1},disabled:{type:Boolean,default:void 0},round:{type:Boolean,default:!0},"onUpdate:value":[Function,Array],onUpdateValue:[Function,Array],checkedValue:{type:[String,Number,Boolean],default:!0},uncheckedValue:{type:[String,Number,Boolean],default:!1},railStyle:Function,rubberBand:{type:Boolean,default:!0},spinProps:Object,onChange:[Function,Array]}),j,M=c({name:`Switch`,props:A,slots:Object,setup(t){j===void 0&&(j=typeof CSS<`u`?CSS.supports===void 0?!1:CSS.supports(`width`,`max(1px)`):!0);let{mergedClsPrefixRef:n,inlineThemeDisabled:a,mergedComponentPropsRef:o}=y(t),c=h(`Switch`,`-switch`,k,O,t,n),l=b(t,{mergedSize(e){return t.size===void 0?e?e.mergedSize.value:o?.value?.Switch?.size||`medium`:t.size}}),{mergedSizeRef:u,mergedDisabledRef:d}=l,p=x(t.defaultValue),m=S(v(t,`value`),p),g=i(()=>m.value===t.checkedValue),C=x(!1),w=x(!1),T=i(()=>{let{railStyle:e}=t;if(e)return e({focused:w.value,checked:g.value})});function E(e){let{"onUpdate:value":n,onChange:i,onUpdateValue:a}=t,{nTriggerFormInput:o,nTriggerFormChange:s}=l;n&&r(n,e),a&&r(a,e),i&&r(i,e),p.value=e,o(),s()}function D(){let{nTriggerFormFocus:e}=l;e()}function A(){let{nTriggerFormBlur:e}=l;e()}function M(){t.loading||d.value||(m.value===t.checkedValue?E(t.uncheckedValue):E(t.checkedValue))}function N(){w.value=!0,D()}function P(){w.value=!1,A(),C.value=!1}function F(e){t.loading||d.value||e.key===` `&&(m.value===t.checkedValue?E(t.uncheckedValue):E(t.checkedValue),C.value=!1)}function I(e){t.loading||d.value||e.key===` `&&(e.preventDefault(),C.value=!0)}let L=i(()=>{let{value:e}=u,{self:{opacityDisabled:t,railColor:n,railColorActive:r,buttonBoxShadow:i,buttonColor:a,boxShadowFocus:o,loadingColor:l,textColor:d,iconColor:p,[s(`buttonHeight`,e)]:m,[s(`buttonWidth`,e)]:h,[s(`buttonWidthPressed`,e)]:g,[s(`railHeight`,e)]:v,[s(`railWidth`,e)]:y,[s(`railBorderRadius`,e)]:b,[s(`buttonBorderRadius`,e)]:x},common:{cubicBezierEaseInOut:S}}=c.value,C,w,T;return j?(C=`calc((${v} - ${m}) / 2)`,w=`max(${v}, ${m})`,T=`max(${y}, calc(${y} + ${m} - ${v}))`):(C=f((_(v)-_(m))/2),w=f(Math.max(_(v),_(m))),T=_(v)>_(m)?y:f(_(y)+_(m)-_(v))),{"--n-bezier":S,"--n-button-border-radius":x,"--n-button-box-shadow":i,"--n-button-color":a,"--n-button-width":h,"--n-button-width-pressed":g,"--n-button-height":m,"--n-height":w,"--n-offset":C,"--n-opacity-disabled":t,"--n-rail-border-radius":b,"--n-rail-color":n,"--n-rail-color-active":r,"--n-rail-height":v,"--n-rail-width":y,"--n-width":T,"--n-box-shadow-focus":o,"--n-loading-color":l,"--n-text-color":d,"--n-icon-color":p}}),R=a?e(`switch`,i(()=>u.value[0]),L,t):void 0;return{handleClick:M,handleBlur:P,handleFocus:N,handleKeyup:F,handleKeydown:I,mergedRailStyle:T,pressed:C,mergedClsPrefix:n,mergedValue:m,checked:g,mergedDisabled:d,cssVars:a?void 0:L,themeClass:R?.themeClass,onRender:R?.onRender}},render(){let{mergedClsPrefix:e,mergedDisabled:t,checked:r,mergedRailStyle:i,onRender:o,$slots:s}=this;o?.();let{checked:c,unchecked:u,icon:f,"checked-icon":p,"unchecked-icon":h}=s,g=!(a(f)&&a(p)&&a(h));return l(`div`,{role:`switch`,"aria-checked":r,class:[`${e}-switch`,this.themeClass,g&&`${e}-switch--icon`,r&&`${e}-switch--active`,t&&`${e}-switch--disabled`,this.round&&`${e}-switch--round`,this.loading&&`${e}-switch--loading`,this.pressed&&`${e}-switch--pressed`,this.rubberBand&&`${e}-switch--rubber-band`],tabindex:this.mergedDisabled?void 0:0,style:this.cssVars,onClick:this.handleClick,onFocus:this.handleFocus,onBlur:this.handleBlur,onKeyup:this.handleKeyup,onKeydown:this.handleKeydown},l(`div`,{class:`${e}-switch__rail`,"aria-hidden":`true`,style:i},n(c,t=>n(u,n=>t||n?l(`div`,{"aria-hidden":!0,class:`${e}-switch__children-placeholder`},l(`div`,{class:`${e}-switch__rail-placeholder`},l(`div`,{class:`${e}-switch__button-placeholder`}),t),l(`div`,{class:`${e}-switch__rail-placeholder`},l(`div`,{class:`${e}-switch__button-placeholder`}),n)):null)),l(`div`,{class:`${e}-switch__button`},n(f,t=>n(p,r=>n(h,n=>l(m,null,{default:()=>this.loading?l(d,Object.assign({key:`loading`,clsPrefix:e,strokeWidth:20},this.spinProps)):this.checked&&(r||t)?l(`div`,{class:`${e}-switch__button-icon`,key:r?`checked-icon`:`icon`},r||t):!this.checked&&(n||t)?l(`div`,{class:`${e}-switch__button-icon`,key:n?`unchecked-icon`:`icon`},n||t):null})))),n(c,t=>t&&l(`div`,{key:`checked`,class:`${e}-switch__checked`},t)),n(u,t=>t&&l(`div`,{key:`unchecked`,class:`${e}-switch__unchecked`},t)))))}});export{M as t};