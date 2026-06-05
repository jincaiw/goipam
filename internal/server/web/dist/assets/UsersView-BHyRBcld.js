import{t as e}from"./Add-9gMbK7KI.js";import{t}from"./Popconfirm-BhS_nm0j.js";import{t as n}from"./Switch-C_WiCcnu.js";import{$r as r,An as i,At as a,Br as o,Cr as s,Ct as c,Dt as l,Fn as u,Fr as d,Gr as f,Hn as p,Hr as m,In as h,Ir as g,Jr as _,Kr as v,Ln as y,Lt as b,Mr as x,Nr as S,O as C,Or as w,Pr as T,Pt as E,Q as ee,Rr as D,Sr as O,St as k,Tr as A,Tt as j,Un as M,Ur as N,Vn as P,Vr as F,Wr as I,Xn as te,Xr as L,Xt as R,Yn as z,Zn as ne,Zr as re,_ as ie,_t as ae,ai as B,ar as oe,br as V,bt as H,cr as se,dn as ce,ei as le,en as U,gn as ue,gr as W,hn as de,hr as fe,ii as G,jn as pe,jt as K,ln as q,lt as me,mn as J,ni as he,qn as Y,qr as ge,ri as X,rr as _e,si as ve,sr as ye,ti as Z,wr as be,wt as xe,xr as Q,xt as Se}from"./index-Dmrq63M5.js";var Ce=z(`.v-x-scroll`,{overflow:`auto`,scrollbarWidth:`none`},[z(`&::-webkit-scrollbar`,{width:0,height:0})]),we=N({name:`XScroll`,props:{disabled:Boolean,onScroll:Function},setup(){let e=X(null);function t(e){!(e.currentTarget.offsetWidth<e.currentTarget.scrollWidth)||e.deltaY===0||(e.currentTarget.scrollLeft+=e.deltaY+e.deltaX,e.preventDefault())}let n=ne();return Ce.mount({id:`vueuc/x-scroll`,head:!0,anchorMetaName:te,ssr:n}),Object.assign({selfRef:e,handleWheel:t},{scrollTo(...t){var n;(n=e.value)==null||n.scrollTo(...t)}})},render(){return I(`div`,{ref:`selfRef`,onScroll:this.onScroll,onWheel:this.disabled?void 0:this.handleWheel,class:`v-x-scroll`},this.$slots)}}),Te=/\s/;function Ee(e){for(var t=e.length;t--&&Te.test(e.charAt(t)););return t}var De=/^\s+/;function Oe(e){return e&&e.slice(0,Ee(e)+1).replace(De,``)}var ke=NaN,Ae=/^[-+]0x[0-9a-f]+$/i,je=/^0b[01]+$/i,Me=/^0o[0-7]+$/i,Ne=parseInt;function Pe(e){if(typeof e==`number`)return e;if(de(e))return ke;if(J(e)){var t=typeof e.valueOf==`function`?e.valueOf():e;e=J(t)?t+``:t}if(typeof e!=`string`)return e===0?e:+e;e=Oe(e);var n=je.test(e);return n||Me.test(e)?Ne(e.slice(2),n?2:8):Ae.test(e)?ke:+e}var Fe=function(){return ue.Date.now()},Ie=`Expected a function`,Le=Math.max,Re=Math.min;function ze(e,t,n){var r,i,a,o,s,c,l=0,u=!1,d=!1,f=!0;if(typeof e!=`function`)throw TypeError(Ie);t=Pe(t)||0,J(n)&&(u=!!n.leading,d=`maxWait`in n,a=d?Le(Pe(n.maxWait)||0,t):a,f=`trailing`in n?!!n.trailing:f);function p(t){var n=r,a=i;return r=i=void 0,l=t,o=e.apply(a,n),o}function m(e){return l=e,s=setTimeout(_,t),u?p(e):o}function h(e){var n=e-c,r=e-l,i=t-n;return d?Re(i,a-r):i}function g(e){var n=e-c,r=e-l;return c===void 0||n>=t||n<0||d&&r>=a}function _(){var e=Fe();if(g(e))return v(e);s=setTimeout(_,h(e))}function v(e){return s=void 0,f&&r?p(e):(r=i=void 0,o)}function y(){s!==void 0&&clearTimeout(s),l=0,r=c=i=s=void 0}function b(){return s===void 0?o:v(Fe())}function x(){var e=Fe(),n=g(e);if(r=arguments,i=this,c=e,n){if(s===void 0)return m(c);if(d)return clearTimeout(s),s=setTimeout(_,t),p(c)}return s===void 0&&(s=setTimeout(_,t)),o}return x.cancel=y,x.flush=b,x}var Be=`Expected a function`;function Ve(e,t,n){var r=!0,i=!0;if(typeof e!=`function`)throw TypeError(Be);return J(n)&&(r=`leading`in n?!!n.leading:r,i=`trailing`in n?!!n.trailing:i),ze(e,t,{leading:r,maxWait:t,trailing:i})}var He={tabFontSizeSmall:`14px`,tabFontSizeMedium:`14px`,tabFontSizeLarge:`16px`,tabGapSmallLine:`36px`,tabGapMediumLine:`36px`,tabGapLargeLine:`36px`,tabGapSmallLineVertical:`8px`,tabGapMediumLineVertical:`8px`,tabGapLargeLineVertical:`8px`,tabPaddingSmallLine:`6px 0`,tabPaddingMediumLine:`10px 0`,tabPaddingLargeLine:`14px 0`,tabPaddingVerticalSmallLine:`6px 12px`,tabPaddingVerticalMediumLine:`8px 16px`,tabPaddingVerticalLargeLine:`10px 20px`,tabGapSmallBar:`36px`,tabGapMediumBar:`36px`,tabGapLargeBar:`36px`,tabGapSmallBarVertical:`8px`,tabGapMediumBarVertical:`8px`,tabGapLargeBarVertical:`8px`,tabPaddingSmallBar:`4px 0`,tabPaddingMediumBar:`6px 0`,tabPaddingLargeBar:`10px 0`,tabPaddingVerticalSmallBar:`6px 12px`,tabPaddingVerticalMediumBar:`8px 16px`,tabPaddingVerticalLargeBar:`10px 20px`,tabGapSmallCard:`4px`,tabGapMediumCard:`4px`,tabGapLargeCard:`4px`,tabGapSmallCardVertical:`4px`,tabGapMediumCardVertical:`4px`,tabGapLargeCardVertical:`4px`,tabPaddingSmallCard:`8px 16px`,tabPaddingMediumCard:`10px 20px`,tabPaddingLargeCard:`12px 24px`,tabPaddingSmallSegment:`4px 0`,tabPaddingMediumSegment:`6px 0`,tabPaddingLargeSegment:`8px 0`,tabPaddingVerticalLargeSegment:`0 8px`,tabPaddingVerticalSmallCard:`8px 12px`,tabPaddingVerticalMediumCard:`10px 16px`,tabPaddingVerticalLargeCard:`12px 20px`,tabPaddingVerticalSmallSegment:`0 4px`,tabPaddingVerticalMediumSegment:`0 6px`,tabGapSmallSegment:`0`,tabGapMediumSegment:`0`,tabGapLargeSegment:`0`,tabGapSmallSegmentVertical:`0`,tabGapMediumSegmentVertical:`0`,tabGapLargeSegmentVertical:`0`,panePaddingSmall:`8px 0 0 0`,panePaddingMedium:`12px 0 0 0`,panePaddingLarge:`16px 0 0 0`,closeSize:`18px`,closeIconSize:`14px`};function Ue(e){let{textColor2:t,primaryColor:n,textColorDisabled:r,closeIconColor:i,closeIconColorHover:a,closeIconColorPressed:o,closeColorHover:s,closeColorPressed:c,tabColor:l,baseColor:u,dividerColor:d,fontWeight:f,textColor1:p,borderRadius:m,fontSize:h,fontWeightStrong:g}=e;return Object.assign(Object.assign({},He),{colorSegment:l,tabFontSizeCard:h,tabTextColorLine:p,tabTextColorActiveLine:n,tabTextColorHoverLine:n,tabTextColorDisabledLine:r,tabTextColorSegment:p,tabTextColorActiveSegment:t,tabTextColorHoverSegment:t,tabTextColorDisabledSegment:r,tabTextColorBar:p,tabTextColorActiveBar:n,tabTextColorHoverBar:n,tabTextColorDisabledBar:r,tabTextColorCard:p,tabTextColorHoverCard:p,tabTextColorActiveCard:n,tabTextColorDisabledCard:r,barColor:n,closeIconColor:i,closeIconColorHover:a,closeIconColorPressed:o,closeColorHover:s,closeColorPressed:c,closeBorderRadius:m,tabColor:l,tabColorSegment:u,tabBorderColor:d,tabFontWeightActive:f,tabFontWeight:f,tabBorderRadius:m,paneTextColor:t,fontWeightStrong:g})}var We={name:`Tabs`,common:R,self:Ue},Ge=_e(`n-tabs`),Ke={tab:[String,Number,Object,Function],name:{type:[String,Number],required:!0},disabled:Boolean,displayDirective:{type:String,default:`if`},closable:{type:Boolean,default:void 0},tabProps:Object,label:[String,Number,Object,Function]},qe=N({__TAB_PANE__:!0,name:`TabPane`,alias:[`TabPanel`],props:Ke,slots:Object,setup(e){let t=f(Ge,null);return t||M(`tab-pane`,"`n-tab-pane` must be placed inside `n-tabs`."),{style:t.paneStyleRef,class:t.paneClassRef,mergedClsPrefix:t.mergedClsPrefixRef}},render(){return I(`div`,{class:[`${this.mergedClsPrefix}-tab-pane`,this.class],style:this.style},this.$slots)}}),Je=N({__TAB__:!0,inheritAttrs:!1,name:`Tab`,props:Object.assign({internalLeftPadded:Boolean,internalAddable:Boolean,internalCreatedByPane:Boolean},y(Ke,[`displayDirective`])),setup(e){let{mergedClsPrefixRef:t,valueRef:n,typeRef:r,closableRef:i,tabStyleRef:a,addTabStyleRef:o,tabClassRef:s,addTabClassRef:c,tabChangeIdRef:l,onBeforeLeaveRef:u,triggerRef:d,handleAdd:p,activateTab:m,handleClose:h}=f(Ge);return{trigger:d,mergedClosable:g(()=>{if(e.internalAddable)return!1;let{closable:t}=e;return t===void 0?i.value:t}),style:a,addStyle:o,tabClass:s,addTabClass:c,clsPrefix:t,value:n,type:r,handleClose(t){t.stopPropagation(),!e.disabled&&h(e.name)},activateTab(){if(e.disabled)return;if(e.internalAddable){p();return}let{name:t}=e,r=++l.id;if(t!==n.value){let{value:i}=u;i?Promise.resolve(i(e.name,n.value)).then(e=>{e&&l.id===r&&m(t)}):m(t)}}}},render(){let{internalAddable:t,clsPrefix:n,name:r,disabled:i,label:a,tab:o,value:s,mergedClosable:c,trigger:l,$slots:{default:u}}=this,d=a??o;return I(`div`,{class:`${n}-tabs-tab-wrapper`},this.internalLeftPadded?I(`div`,{class:`${n}-tabs-tab-pad`}):null,I(`div`,Object.assign({key:r,"data-name":r,"data-disabled":i?!0:void 0},v({class:[`${n}-tabs-tab`,s===r&&`${n}-tabs-tab--active`,i&&`${n}-tabs-tab--disabled`,c&&`${n}-tabs-tab--closable`,t&&`${n}-tabs-tab--addable`,t?this.addTabClass:this.tabClass],onClick:l===`click`?this.activateTab:void 0,onMouseenter:l===`hover`?this.activateTab:void 0,style:t?this.addStyle:this.style},this.internalCreatedByPane?this.tabProps||{}:this.$attrs)),I(`span`,{class:`${n}-tabs-tab__label`},t?I(T,null,I(`div`,{class:`${n}-tabs-tab__height-placeholder`},`\xA0`),I(q,{clsPrefix:n},{default:()=>I(e,null)})):u?u():typeof d==`object`?d:h(d??r)),c&&this.type===`card`?I(U,{clsPrefix:n,class:`${n}-tabs-tab__close`,onClick:this.handleClose,disabled:i}):null))}}),Ye=Q(`tabs`,`
 box-sizing: border-box;
 width: 100%;
 display: flex;
 flex-direction: column;
 transition:
 background-color .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
`,[s(`segment-type`,[Q(`tabs-rail`,[V(`&.transition-disabled`,[Q(`tabs-capsule`,`
 transition: none;
 `)])])]),s(`top`,[Q(`tab-pane`,`
 padding: var(--n-pane-padding-top) var(--n-pane-padding-right) var(--n-pane-padding-bottom) var(--n-pane-padding-left);
 `)]),s(`left`,[Q(`tab-pane`,`
 padding: var(--n-pane-padding-right) var(--n-pane-padding-bottom) var(--n-pane-padding-left) var(--n-pane-padding-top);
 `)]),s(`left, right`,`
 flex-direction: row;
 `,[Q(`tabs-bar`,`
 width: 2px;
 right: 0;
 transition:
 top .2s var(--n-bezier),
 max-height .2s var(--n-bezier),
 background-color .3s var(--n-bezier);
 `),Q(`tabs-tab`,`
 padding: var(--n-tab-padding-vertical); 
 `)]),s(`right`,`
 flex-direction: row-reverse;
 `,[Q(`tab-pane`,`
 padding: var(--n-pane-padding-left) var(--n-pane-padding-top) var(--n-pane-padding-right) var(--n-pane-padding-bottom);
 `),Q(`tabs-bar`,`
 left: 0;
 `)]),s(`bottom`,`
 flex-direction: column-reverse;
 justify-content: flex-end;
 `,[Q(`tab-pane`,`
 padding: var(--n-pane-padding-bottom) var(--n-pane-padding-right) var(--n-pane-padding-top) var(--n-pane-padding-left);
 `),Q(`tabs-bar`,`
 top: 0;
 `)]),Q(`tabs-rail`,`
 position: relative;
 padding: 3px;
 border-radius: var(--n-tab-border-radius);
 width: 100%;
 background-color: var(--n-color-segment);
 transition: background-color .3s var(--n-bezier);
 display: flex;
 align-items: center;
 `,[Q(`tabs-capsule`,`
 border-radius: var(--n-tab-border-radius);
 position: absolute;
 pointer-events: none;
 background-color: var(--n-tab-color-segment);
 box-shadow: 0 1px 3px 0 rgba(0, 0, 0, .08);
 transition: transform 0.3s var(--n-bezier);
 `),Q(`tabs-tab-wrapper`,`
 flex-basis: 0;
 flex-grow: 1;
 display: flex;
 align-items: center;
 justify-content: center;
 `,[Q(`tabs-tab`,`
 overflow: hidden;
 border-radius: var(--n-tab-border-radius);
 width: 100%;
 display: flex;
 align-items: center;
 justify-content: center;
 `,[s(`active`,`
 font-weight: var(--n-font-weight-strong);
 color: var(--n-tab-text-color-active);
 `),V(`&:hover`,`
 color: var(--n-tab-text-color-hover);
 `)])])]),s(`flex`,[Q(`tabs-nav`,`
 width: 100%;
 position: relative;
 `,[Q(`tabs-wrapper`,`
 width: 100%;
 `,[Q(`tabs-tab`,`
 margin-right: 0;
 `)])])]),Q(`tabs-nav`,`
 box-sizing: border-box;
 line-height: 1.5;
 display: flex;
 transition: border-color .3s var(--n-bezier);
 `,[O(`prefix, suffix`,`
 display: flex;
 align-items: center;
 `),O(`prefix`,`padding-right: 16px;`),O(`suffix`,`padding-left: 16px;`)]),s(`top, bottom`,[V(`>`,[Q(`tabs-nav`,[Q(`tabs-nav-scroll-wrapper`,[V(`&::before`,`
 top: 0;
 bottom: 0;
 left: 0;
 width: 20px;
 `),V(`&::after`,`
 top: 0;
 bottom: 0;
 right: 0;
 width: 20px;
 `),s(`shadow-start`,[V(`&::before`,`
 box-shadow: inset 10px 0 8px -8px rgba(0, 0, 0, .12);
 `)]),s(`shadow-end`,[V(`&::after`,`
 box-shadow: inset -10px 0 8px -8px rgba(0, 0, 0, .12);
 `)])])])])]),s(`left, right`,[Q(`tabs-nav-scroll-content`,`
 flex-direction: column;
 `),V(`>`,[Q(`tabs-nav`,[Q(`tabs-nav-scroll-wrapper`,[V(`&::before`,`
 top: 0;
 left: 0;
 right: 0;
 height: 20px;
 `),V(`&::after`,`
 bottom: 0;
 left: 0;
 right: 0;
 height: 20px;
 `),s(`shadow-start`,[V(`&::before`,`
 box-shadow: inset 0 10px 8px -8px rgba(0, 0, 0, .12);
 `)]),s(`shadow-end`,[V(`&::after`,`
 box-shadow: inset 0 -10px 8px -8px rgba(0, 0, 0, .12);
 `)])])])])]),Q(`tabs-nav-scroll-wrapper`,`
 flex: 1;
 position: relative;
 overflow: hidden;
 `,[Q(`tabs-nav-y-scroll`,`
 height: 100%;
 width: 100%;
 overflow-y: auto; 
 scrollbar-width: none;
 `,[V(`&::-webkit-scrollbar, &::-webkit-scrollbar-track-piece, &::-webkit-scrollbar-thumb`,`
 width: 0;
 height: 0;
 display: none;
 `)]),V(`&::before, &::after`,`
 transition: box-shadow .3s var(--n-bezier);
 pointer-events: none;
 content: "";
 position: absolute;
 z-index: 1;
 `)]),Q(`tabs-nav-scroll-content`,`
 display: flex;
 position: relative;
 min-width: 100%;
 min-height: 100%;
 width: fit-content;
 box-sizing: border-box;
 `),Q(`tabs-wrapper`,`
 display: inline-flex;
 flex-wrap: nowrap;
 position: relative;
 `),Q(`tabs-tab-wrapper`,`
 display: flex;
 flex-wrap: nowrap;
 flex-shrink: 0;
 flex-grow: 0;
 `),Q(`tabs-tab`,`
 cursor: pointer;
 white-space: nowrap;
 flex-wrap: nowrap;
 display: inline-flex;
 align-items: center;
 color: var(--n-tab-text-color);
 font-size: var(--n-tab-font-size);
 background-clip: padding-box;
 padding: var(--n-tab-padding);
 transition:
 box-shadow .3s var(--n-bezier),
 color .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
 `,[s(`disabled`,{cursor:`not-allowed`}),O(`close`,`
 margin-left: 6px;
 transition:
 background-color .3s var(--n-bezier),
 color .3s var(--n-bezier);
 `),O(`label`,`
 display: flex;
 align-items: center;
 z-index: 1;
 `)]),Q(`tabs-bar`,`
 position: absolute;
 bottom: 0;
 height: 2px;
 border-radius: 1px;
 background-color: var(--n-bar-color);
 transition:
 left .2s var(--n-bezier),
 max-width .2s var(--n-bezier),
 opacity .3s var(--n-bezier),
 background-color .3s var(--n-bezier);
 `,[V(`&.transition-disabled`,`
 transition: none;
 `),s(`disabled`,`
 background-color: var(--n-tab-text-color-disabled)
 `)]),Q(`tabs-pane-wrapper`,`
 position: relative;
 overflow: hidden;
 transition: max-height .2s var(--n-bezier);
 `),Q(`tab-pane`,`
 color: var(--n-pane-text-color);
 width: 100%;
 transition:
 color .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 opacity .2s var(--n-bezier);
 left: 0;
 right: 0;
 top: 0;
 `,[V(`&.next-transition-leave-active, &.prev-transition-leave-active, &.next-transition-enter-active, &.prev-transition-enter-active`,`
 transition:
 color .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 transform .2s var(--n-bezier),
 opacity .2s var(--n-bezier);
 `),V(`&.next-transition-leave-active, &.prev-transition-leave-active`,`
 position: absolute;
 `),V(`&.next-transition-enter-from, &.prev-transition-leave-to`,`
 transform: translateX(32px);
 opacity: 0;
 `),V(`&.next-transition-leave-to, &.prev-transition-enter-from`,`
 transform: translateX(-32px);
 opacity: 0;
 `),V(`&.next-transition-leave-from, &.next-transition-enter-to, &.prev-transition-leave-from, &.prev-transition-enter-to`,`
 transform: translateX(0);
 opacity: 1;
 `)]),Q(`tabs-tab-pad`,`
 box-sizing: border-box;
 width: var(--n-tab-gap);
 flex-grow: 0;
 flex-shrink: 0;
 `),s(`line-type, bar-type`,[Q(`tabs-tab`,`
 font-weight: var(--n-tab-font-weight);
 box-sizing: border-box;
 vertical-align: bottom;
 `,[V(`&:hover`,{color:`var(--n-tab-text-color-hover)`}),s(`active`,`
 color: var(--n-tab-text-color-active);
 font-weight: var(--n-tab-font-weight-active);
 `),s(`disabled`,{color:`var(--n-tab-text-color-disabled)`})])]),Q(`tabs-nav`,[s(`line-type`,[s(`top`,[O(`prefix, suffix`,`
 border-bottom: 1px solid var(--n-tab-border-color);
 `),Q(`tabs-nav-scroll-content`,`
 border-bottom: 1px solid var(--n-tab-border-color);
 `),Q(`tabs-bar`,`
 bottom: -1px;
 `)]),s(`left`,[O(`prefix, suffix`,`
 border-right: 1px solid var(--n-tab-border-color);
 `),Q(`tabs-nav-scroll-content`,`
 border-right: 1px solid var(--n-tab-border-color);
 `),Q(`tabs-bar`,`
 right: -1px;
 `)]),s(`right`,[O(`prefix, suffix`,`
 border-left: 1px solid var(--n-tab-border-color);
 `),Q(`tabs-nav-scroll-content`,`
 border-left: 1px solid var(--n-tab-border-color);
 `),Q(`tabs-bar`,`
 left: -1px;
 `)]),s(`bottom`,[O(`prefix, suffix`,`
 border-top: 1px solid var(--n-tab-border-color);
 `),Q(`tabs-nav-scroll-content`,`
 border-top: 1px solid var(--n-tab-border-color);
 `),Q(`tabs-bar`,`
 top: -1px;
 `)]),O(`prefix, suffix`,`
 transition: border-color .3s var(--n-bezier);
 `),Q(`tabs-nav-scroll-content`,`
 transition: border-color .3s var(--n-bezier);
 `),Q(`tabs-bar`,`
 border-radius: 0;
 `)]),s(`card-type`,[O(`prefix, suffix`,`
 transition: border-color .3s var(--n-bezier);
 `),Q(`tabs-pad`,`
 flex-grow: 1;
 transition: border-color .3s var(--n-bezier);
 `),Q(`tabs-tab-pad`,`
 transition: border-color .3s var(--n-bezier);
 `),Q(`tabs-tab`,`
 font-weight: var(--n-tab-font-weight);
 border: 1px solid var(--n-tab-border-color);
 background-color: var(--n-tab-color);
 box-sizing: border-box;
 position: relative;
 vertical-align: bottom;
 display: flex;
 justify-content: space-between;
 font-size: var(--n-tab-font-size);
 color: var(--n-tab-text-color);
 `,[s(`addable`,`
 padding-left: 8px;
 padding-right: 8px;
 font-size: 16px;
 justify-content: center;
 `,[O(`height-placeholder`,`
 width: 0;
 font-size: var(--n-tab-font-size);
 `),be(`disabled`,[V(`&:hover`,`
 color: var(--n-tab-text-color-hover);
 `)])]),s(`closable`,`padding-right: 8px;`),s(`active`,`
 background-color: #0000;
 font-weight: var(--n-tab-font-weight-active);
 color: var(--n-tab-text-color-active);
 `),s(`disabled`,`color: var(--n-tab-text-color-disabled);`)])]),s(`left, right`,`
 flex-direction: column; 
 `,[O(`prefix, suffix`,`
 padding: var(--n-tab-padding-vertical);
 `),Q(`tabs-wrapper`,`
 flex-direction: column;
 `),Q(`tabs-tab-wrapper`,`
 flex-direction: column;
 `,[Q(`tabs-tab-pad`,`
 height: var(--n-tab-gap-vertical);
 width: 100%;
 `)])]),s(`top`,[s(`card-type`,[Q(`tabs-scroll-padding`,`border-bottom: 1px solid var(--n-tab-border-color);`),O(`prefix, suffix`,`
 border-bottom: 1px solid var(--n-tab-border-color);
 `),Q(`tabs-tab`,`
 border-top-left-radius: var(--n-tab-border-radius);
 border-top-right-radius: var(--n-tab-border-radius);
 `,[s(`active`,`
 border-bottom: 1px solid #0000;
 `)]),Q(`tabs-tab-pad`,`
 border-bottom: 1px solid var(--n-tab-border-color);
 `),Q(`tabs-pad`,`
 border-bottom: 1px solid var(--n-tab-border-color);
 `)])]),s(`left`,[s(`card-type`,[Q(`tabs-scroll-padding`,`border-right: 1px solid var(--n-tab-border-color);`),O(`prefix, suffix`,`
 border-right: 1px solid var(--n-tab-border-color);
 `),Q(`tabs-tab`,`
 border-top-left-radius: var(--n-tab-border-radius);
 border-bottom-left-radius: var(--n-tab-border-radius);
 `,[s(`active`,`
 border-right: 1px solid #0000;
 `)]),Q(`tabs-tab-pad`,`
 border-right: 1px solid var(--n-tab-border-color);
 `),Q(`tabs-pad`,`
 border-right: 1px solid var(--n-tab-border-color);
 `)])]),s(`right`,[s(`card-type`,[Q(`tabs-scroll-padding`,`border-left: 1px solid var(--n-tab-border-color);`),O(`prefix, suffix`,`
 border-left: 1px solid var(--n-tab-border-color);
 `),Q(`tabs-tab`,`
 border-top-right-radius: var(--n-tab-border-radius);
 border-bottom-right-radius: var(--n-tab-border-radius);
 `,[s(`active`,`
 border-left: 1px solid #0000;
 `)]),Q(`tabs-tab-pad`,`
 border-left: 1px solid var(--n-tab-border-color);
 `),Q(`tabs-pad`,`
 border-left: 1px solid var(--n-tab-border-color);
 `)])]),s(`bottom`,[s(`card-type`,[Q(`tabs-scroll-padding`,`border-top: 1px solid var(--n-tab-border-color);`),O(`prefix, suffix`,`
 border-top: 1px solid var(--n-tab-border-color);
 `),Q(`tabs-tab`,`
 border-bottom-left-radius: var(--n-tab-border-radius);
 border-bottom-right-radius: var(--n-tab-border-radius);
 `,[s(`active`,`
 border-top: 1px solid #0000;
 `)]),Q(`tabs-tab-pad`,`
 border-top: 1px solid var(--n-tab-border-color);
 `),Q(`tabs-pad`,`
 border-top: 1px solid var(--n-tab-border-color);
 `)])])])]),$=Ve,Xe=N({name:`Tabs`,props:Object.assign(Object.assign({},ce.props),{value:[String,Number],defaultValue:[String,Number],trigger:{type:String,default:`click`},type:{type:String,default:`bar`},closable:Boolean,justifyContent:String,size:String,placement:{type:String,default:`top`},tabStyle:[String,Object],tabClass:String,addTabStyle:[String,Object],addTabClass:String,barWidth:Number,paneClass:String,paneStyle:[String,Object],paneWrapperClass:String,paneWrapperStyle:[String,Object],addable:[Boolean,Object],tabsPadding:{type:Number,default:0},animated:Boolean,onBeforeLeave:Function,onAdd:Function,"onUpdate:value":[Function,Array],onUpdateValue:[Function,Array],onClose:[Function,Array],labelSize:String,activeName:[String,Number],onActiveNameChange:[Function,Array]}),slots:Object,setup(e,{slots:t}){let{mergedClsPrefixRef:n,inlineThemeDisabled:a,mergedComponentPropsRef:o}=pe(e),s=ce(`Tabs`,`-tabs`,Ye,We,e,n),c=X(null),l=X(null),u=X(null),d=X(null),f=X(null),m=X(null),h=X(!0),v=X(!0),y=oe(e,[`labelSize`,`size`]),b=g(()=>y.value?y.value:o?.value?.Tabs?.size||`medium`),x=oe(e,[`activeName`,`value`]),S=X(x.value??e.defaultValue??(t.default?P(t.default())[0]?.props?.name:null)),C=ye(x,S),w={id:0},T=g(()=>{if(!(!e.justifyContent||e.type===`card`))return{display:`flex`,justifyContent:e.justifyContent}});r(C,()=>{w.id=0,k(),j()});function E(){let{value:e}=C;return e===null?null:c.value?.querySelector(`[data-name="${e}"]`)}function ee(t){if(e.type===`card`)return;let{value:r}=l;if(!r)return;let i=r.style.opacity===`0`;if(t){let a=`${n.value}-tabs-bar--disabled`,{barWidth:o,placement:s}=e;if(t.dataset.disabled===`true`?r.classList.add(a):r.classList.remove(a),[`top`,`bottom`].includes(s)){if(O([`top`,`maxHeight`,`height`]),typeof o==`number`&&t.offsetWidth>=o){let e=Math.floor((t.offsetWidth-o)/2)+t.offsetLeft;r.style.left=`${e}px`,r.style.maxWidth=`${o}px`}else r.style.left=`${t.offsetLeft}px`,r.style.maxWidth=`${t.offsetWidth}px`;r.style.width=`8192px`,i&&(r.style.transition=`none`),r.offsetWidth,i&&(r.style.transition=``,r.style.opacity=`1`)}else{if(O([`left`,`maxWidth`,`width`]),typeof o==`number`&&t.offsetHeight>=o){let e=Math.floor((t.offsetHeight-o)/2)+t.offsetTop;r.style.top=`${e}px`,r.style.maxHeight=`${o}px`}else r.style.top=`${t.offsetTop}px`,r.style.maxHeight=`${t.offsetHeight}px`;r.style.height=`8192px`,i&&(r.style.transition=`none`),r.offsetHeight,i&&(r.style.transition=``,r.style.opacity=`1`)}}}function D(){if(e.type===`card`)return;let{value:t}=l;t&&(t.style.opacity=`0`)}function O(e){let{value:t}=l;if(t)for(let n of e)t.style[n]=``}function k(){if(e.type===`card`)return;let t=E();t?ee(t):D()}function j(){let e=f.value?.$el;if(!e)return;let t=E();if(!t)return;let{scrollLeft:n,offsetWidth:r}=e,{offsetLeft:i,offsetWidth:a}=t;n>i?e.scrollTo({top:0,left:i,behavior:`smooth`}):i+a>n+r&&e.scrollTo({top:0,left:i+a-r,behavior:`smooth`})}let M=X(null),N=0,F=null;function I(e){let t=M.value;if(t){N=e.getBoundingClientRect().height;let n=`${N}px`,r=()=>{t.style.height=n,t.style.maxHeight=n};F?(r(),F(),F=null):F=r}}function te(e){let t=M.value;if(t){let n=e.getBoundingClientRect().height,r=()=>{document.body.offsetHeight,t.style.maxHeight=`${n}px`,t.style.height=`${Math.max(N,n)}px`};F?(F(),F=null,r()):F=r}}function L(){let t=M.value;if(t){t.style.maxHeight=``,t.style.height=``;let{paneWrapperStyle:n}=e;if(typeof n==`string`)t.style.cssText=n;else if(n){let{maxHeight:e,height:r}=n;e!==void 0&&(t.style.maxHeight=e),r!==void 0&&(t.style.height=r)}}}let R={value:[]},z=X(`next`);function ne(e){let t=C.value,n=`next`;for(let r of R.value){if(r===t)break;if(r===e){n=`prev`;break}}z.value=n,ie(e)}function ie(t){let{onActiveNameChange:n,onUpdateValue:r,"onUpdate:value":i}=e;n&&p(n,t),r&&p(r,t),i&&p(i,t),S.value=t}function ae(t){let{onClose:n}=e;n&&p(n,t)}let B=!0;function V(){let{value:e}=l;if(!e)return;B||=!1;let t=`transition-disabled`;e.classList.add(t),k(),e.classList.remove(t)}let H=X(null);function U({transitionDisabled:e}){let t=c.value;if(!t)return;e&&t.classList.add(`transition-disabled`);let n=E();n&&H.value&&(H.value.style.width=`${n.offsetWidth}px`,H.value.style.height=`${n.offsetHeight}px`,H.value.style.transform=`translateX(${n.offsetLeft-fe(getComputedStyle(t).paddingLeft)}px)`,e&&H.value.offsetWidth),e&&t.classList.remove(`transition-disabled`)}r([C],()=>{e.type===`segment`&&ge(()=>{U({transitionDisabled:!1})})}),_(()=>{e.type===`segment`&&U({transitionDisabled:!0})});let ue=0;function de(t){if(t.contentRect.width===0&&t.contentRect.height===0||ue===t.contentRect.width)return;ue=t.contentRect.width;let{type:n}=e;if((n===`line`||n===`bar`)&&(B||e.justifyContent?.startsWith(`space`))&&V(),n!==`segment`){let{placement:t}=e;Y((t===`top`||t===`bottom`?f.value?.$el:m.value)||null)}}let K=$(de,64);r([()=>e.justifyContent,()=>e.size],()=>{ge(()=>{let{type:t}=e;(t===`line`||t===`bar`)&&V()})});let q=X(!1);function me(t){let{target:n,contentRect:{width:r,height:i}}=t,a=n.parentElement.parentElement.offsetWidth,o=n.parentElement.parentElement.offsetHeight,{placement:s}=e;if(!q.value)s===`top`||s===`bottom`?a<r&&(q.value=!0):o<i&&(q.value=!0);else{let{value:e}=d;if(!e)return;s===`top`||s===`bottom`?a-r>e.$el.offsetWidth&&(q.value=!1):o-i>e.$el.offsetHeight&&(q.value=!1)}Y(f.value?.$el||null)}let J=$(me,64);function he(){let{onAdd:t}=e;t&&t(),ge(()=>{let e=E(),{value:t}=f;!e||!t||t.scrollTo({left:e.offsetLeft,top:0,behavior:`smooth`})})}function Y(t){if(!t)return;let{placement:n}=e;if(n===`top`||n===`bottom`){let{scrollLeft:e,scrollWidth:n,offsetWidth:r}=t;h.value=e<=0,v.value=e+r>=n}else{let{scrollTop:e,scrollHeight:n,offsetHeight:r}=t;h.value=e<=0,v.value=e+r>=n}}let _e=$(e=>{Y(e.target)},64);re(Ge,{triggerRef:G(e,`trigger`),tabStyleRef:G(e,`tabStyle`),tabClassRef:G(e,`tabClass`),addTabStyleRef:G(e,`addTabStyle`),addTabClassRef:G(e,`addTabClass`),paneClassRef:G(e,`paneClass`),paneStyleRef:G(e,`paneStyle`),mergedClsPrefixRef:n,typeRef:G(e,`type`),closableRef:G(e,`closable`),valueRef:C,tabChangeIdRef:w,onBeforeLeaveRef:G(e,`onBeforeLeave`),activateTab:ne,handleClose:ae,handleAdd:he}),se(()=>{k(),j()}),le(()=>{let{value:e}=u;if(!e)return;let{value:t}=n,r=`${t}-tabs-nav-scroll-wrapper--shadow-start`,i=`${t}-tabs-nav-scroll-wrapper--shadow-end`;h.value?e.classList.remove(r):e.classList.add(r),v.value?e.classList.remove(i):e.classList.add(i)});let ve={syncBarPosition:()=>{k()}},Z=()=>{U({transitionDisabled:!0})},be=g(()=>{let{value:t}=b,{type:n}=e,r=`${t}${{card:`Card`,bar:`Bar`,line:`Line`,segment:`Segment`}[n]}`,{self:{barColor:i,closeIconColor:a,closeIconColorHover:o,closeIconColorPressed:c,tabColor:l,tabBorderColor:u,paneTextColor:d,tabFontWeight:f,tabBorderRadius:p,tabFontWeightActive:m,colorSegment:h,fontWeightStrong:g,tabColorSegment:_,closeSize:v,closeIconSize:y,closeColorHover:x,closeColorPressed:S,closeBorderRadius:C,[A(`panePadding`,t)]:w,[A(`tabPadding`,r)]:T,[A(`tabPaddingVertical`,r)]:E,[A(`tabGap`,r)]:ee,[A(`tabGap`,`${r}Vertical`)]:D,[A(`tabTextColor`,n)]:O,[A(`tabTextColorActive`,n)]:k,[A(`tabTextColorHover`,n)]:j,[A(`tabTextColorDisabled`,n)]:M,[A(`tabFontSize`,t)]:N},common:{cubicBezierEaseInOut:P}}=s.value;return{"--n-bezier":P,"--n-color-segment":h,"--n-bar-color":i,"--n-tab-font-size":N,"--n-tab-text-color":O,"--n-tab-text-color-active":k,"--n-tab-text-color-disabled":M,"--n-tab-text-color-hover":j,"--n-pane-text-color":d,"--n-tab-border-color":u,"--n-tab-border-radius":p,"--n-close-size":v,"--n-close-icon-size":y,"--n-close-color-hover":x,"--n-close-color-pressed":S,"--n-close-border-radius":C,"--n-close-icon-color":a,"--n-close-icon-color-hover":o,"--n-close-icon-color-pressed":c,"--n-tab-color":l,"--n-tab-font-weight":f,"--n-tab-font-weight-active":m,"--n-tab-padding":T,"--n-tab-padding-vertical":E,"--n-tab-gap":ee,"--n-tab-gap-vertical":D,"--n-pane-padding-left":W(w,`left`),"--n-pane-padding-right":W(w,`right`),"--n-pane-padding-top":W(w,`top`),"--n-pane-padding-bottom":W(w,`bottom`),"--n-font-weight-strong":g,"--n-tab-color-segment":_}}),xe=a?i(`tabs`,g(()=>`${b.value[0]}${e.type[0]}`),be,e):void 0;return Object.assign({mergedClsPrefix:n,mergedValue:C,renderedNames:new Set,segmentCapsuleElRef:H,tabsPaneWrapperRef:M,tabsElRef:c,barElRef:l,addTabInstRef:d,xScrollInstRef:f,scrollWrapperElRef:u,addTabFixed:q,tabWrapperStyle:T,handleNavResize:K,mergedSize:b,handleScroll:_e,handleTabsResize:J,cssVars:a?void 0:be,themeClass:xe?.themeClass,animationDirection:z,renderNameListRef:R,yScrollElRef:m,handleSegmentResize:Z,onAnimationBeforeLeave:I,onAnimationEnter:te,onAnimationAfterEnter:L,onRender:xe?.onRender},ve)},render(){let{mergedClsPrefix:e,type:t,placement:n,addTabFixed:r,addable:i,mergedSize:a,renderNameListRef:o,onRender:s,paneWrapperClass:c,paneWrapperStyle:l,$slots:{default:d,prefix:f,suffix:p}}=this;s?.();let m=d?P(d()).filter(e=>e.type.__TAB_PANE__===!0):[],h=d?P(d()).filter(e=>e.type.__TAB__===!0):[],g=!h.length,_=t===`card`,v=t===`segment`,y=!_&&!v&&this.justifyContent;o.value=[];let b=()=>{let t=I(`div`,{style:this.tabWrapperStyle,class:`${e}-tabs-wrapper`},y?null:I(`div`,{class:`${e}-tabs-scroll-padding`,style:n===`top`||n===`bottom`?{width:`${this.tabsPadding}px`}:{height:`${this.tabsPadding}px`}}),g?m.map((e,t)=>(o.value.push(e.props.name),et(I(Je,Object.assign({},e.props,{internalCreatedByPane:!0,internalLeftPadded:t!==0&&(!y||y===`center`||y===`start`||y===`end`)}),e.children?{default:e.children.tab}:void 0)))):h.map((e,t)=>(o.value.push(e.props.name),et(t!==0&&!y?$e(e):e))),!r&&i&&_?Qe(i,(g?m.length:h.length)!==0):null,y?null:I(`div`,{class:`${e}-tabs-scroll-padding`,style:{width:`${this.tabsPadding}px`}}));return I(`div`,{ref:`tabsElRef`,class:`${e}-tabs-nav-scroll-content`},_&&i?I(Y,{onResize:this.handleTabsResize},{default:()=>t}):t,_?I(`div`,{class:`${e}-tabs-pad`}):null,_?null:I(`div`,{ref:`barElRef`,class:`${e}-tabs-bar`}))},x=v?`top`:n;return I(`div`,{class:[`${e}-tabs`,this.themeClass,`${e}-tabs--${t}-type`,`${e}-tabs--${a}-size`,y&&`${e}-tabs--flex`,`${e}-tabs--${x}`],style:this.cssVars},I(`div`,{class:[`${e}-tabs-nav--${t}-type`,`${e}-tabs-nav--${x}`,`${e}-tabs-nav`]},u(f,t=>t&&I(`div`,{class:`${e}-tabs-nav__prefix`},t)),v?I(Y,{onResize:this.handleSegmentResize},{default:()=>I(`div`,{class:`${e}-tabs-rail`,ref:`tabsElRef`},I(`div`,{class:`${e}-tabs-capsule`,ref:`segmentCapsuleElRef`},I(`div`,{class:`${e}-tabs-wrapper`},I(`div`,{class:`${e}-tabs-tab`}))),g?m.map((e,t)=>(o.value.push(e.props.name),I(Je,Object.assign({},e.props,{internalCreatedByPane:!0,internalLeftPadded:t!==0}),e.children?{default:e.children.tab}:void 0))):h.map((e,t)=>(o.value.push(e.props.name),t===0?e:$e(e))))}):I(Y,{onResize:this.handleNavResize},{default:()=>I(`div`,{class:`${e}-tabs-nav-scroll-wrapper`,ref:`scrollWrapperElRef`},[`top`,`bottom`].includes(x)?I(we,{ref:`xScrollInstRef`,onScroll:this.handleScroll},{default:b}):I(`div`,{class:`${e}-tabs-nav-y-scroll`,onScroll:this.handleScroll,ref:`yScrollElRef`},b()))}),r&&i&&_?Qe(i,!0):null,u(p,t=>t&&I(`div`,{class:`${e}-tabs-nav__suffix`},t))),g&&(this.animated&&(x===`top`||x===`bottom`)?I(`div`,{ref:`tabsPaneWrapperRef`,style:l,class:[`${e}-tabs-pane-wrapper`,c]},Ze(m,this.mergedValue,this.renderedNames,this.onAnimationBeforeLeave,this.onAnimationEnter,this.onAnimationAfterEnter,this.animationDirection)):Ze(m,this.mergedValue,this.renderedNames)))}});function Ze(e,t,n,r,i,a,o){let s=[];return e.forEach(e=>{let{name:r,displayDirective:i,"display-directive":a}=e.props,o=e=>i===e||a===e,c=t===r;if(e.key!==void 0&&(e.key=r),c||o(`show`)||o(`show:lazy`)&&n.has(r)){n.has(r)||n.add(r);let t=!o(`if`);s.push(t?he(e,[[S,c]]):e)}}),o?I(x,{name:`${o}-transition`,onBeforeLeave:r,onEnter:i,onAfterEnter:a},{default:()=>s}):s}function Qe(e,t){return I(Je,{ref:`addTabInstRef`,key:`__addable`,name:`__addable`,internalCreatedByPane:!0,internalAddable:!0,internalLeftPadded:t,disabled:typeof e==`object`&&e.disabled})}function $e(e){let t=d(e);return t.props?t.props.internalLeftPadded=!0:t.props={internalLeftPadded:!0},t}function et(e){return Array.isArray(e.dynamicProps)?e.dynamicProps.includes(`internalLeftPadded`)||e.dynamicProps.push(`internalLeftPadded`):e.dynamicProps=[`internalLeftPadded`],e}var tt=N({__name:`UsersView`,setup(e){let r=c(),{t:i}=w(),s=X(!1),u=X([]),d=X({page:1,pageSize:20,itemCount:0}),f=X(!1),p=X(null),h=X(!1),v=X(null),y=X(`users`),x=[{label:`Admin`,value:`admin`},{label:`Operator`,value:`operator`},{label:`Viewer`,value:`viewer`}],S=X({username:``,password:``,email:``,is_active:!0,is_superuser:!1,roles:[]}),T={username:{required:!0,message:()=>i(`common.required`),trigger:`blur`},password:{required:!0,message:()=>i(`common.required`),trigger:`blur`,validator(e,t){return!v.value&&!t?Error(i(`common.required`)):!0}}},O=g(()=>[{title:i(`common.id`),key:`id`,width:80},{title:i(`user.username`),key:`username`,width:140},{title:i(`user.email`),key:`email`,ellipsis:{tooltip:!0},render(e){return e.email||`-`}},{title:i(`user.role`),key:`roles`,width:200,render(e){return e.is_superuser?I(b,{type:`error`,size:`small`},{default:()=>`Superuser`}):I(k,{size:4},{default:()=>(e.roles||[]).map(e=>I(b,{type:`info`,size:`small`},{default:()=>e}))})}},{title:i(`common.status`),key:`is_active`,width:100,render(e){return I(b,{type:e.is_active?`success`:`error`,size:`small`},{default:()=>e.is_active?i(`user.active`):i(`user.inactive`)})}},{title:i(`common.actions`),key:`actions`,width:160,render(e){return I(k,{size:`small`},{default:()=>[I(K,{size:`small`,quaternary:!0,onClick:()=>z(e)},{default:()=>i(`common.edit`)}),I(t,{onPositiveClick:()=>re(e.id)},{trigger:()=>I(K,{size:`small`,type:`error`,quaternary:!0},{default:()=>i(`common.delete`)}),default:()=>i(`common.deleteConfirm`)})]})}}]),A=g(()=>[{title:i(`user.role`),key:`name`,width:160},{title:i(`common.description`),key:`description`}]),M=[{name:`admin`,description:`Full access to all resources including user management`},{name:`operator`,description:`Can create, edit, and delete IPAM resources`},{name:`viewer`,description:`Read-only access to all resources`}];async function N(){s.value=!0;try{let e=await ee(d.value.page,d.value.pageSize);e.success&&e.data&&(u.value=e.data.items,d.value.itemCount=e.data.total)}catch(e){r.error(e?.response?.data?.message||i(`user.loadFailed`))}finally{s.value=!1}}function P(e){d.value.page=e,N()}function te(e){d.value.pageSize=e,d.value.page=1,N()}function R(){v.value=null,S.value={username:``,password:``,email:``,is_active:!0,is_superuser:!1,roles:[]},f.value=!0}function z(e){v.value=e.id,S.value={username:e.username,password:``,email:e.email,is_active:e.is_active,is_superuser:e.is_superuser,roles:e.roles||[]},f.value=!0}async function ne(){try{await p.value?.validate()}catch{return}h.value=!0;try{let e={username:S.value.username,email:S.value.email||void 0,is_active:S.value.is_active,is_superuser:S.value.is_superuser,roles:S.value.roles};(!v.value||S.value.password)&&(e.password=S.value.password);let t;t=v.value?await me(v.value,e):await ie(e),t.success?(r.success(v.value?i(`user.updated`):i(`user.created`)),f.value=!1,N()):r.error(t.message||i(`user.saveFailed`))}catch(e){r.error(e?.response?.data?.message||i(`user.saveFailed`))}finally{h.value=!1}}async function re(e){try{let t=await C(e);t.success?(r.success(i(`user.deleted`)),N()):r.error(t.message||i(`user.deleteFailed`))}catch(e){r.error(e?.response?.data?.message||i(`user.deleteFailed`))}}return _(N),(e,t)=>(L(),o(`div`,null,[m(B(ae),{title:B(i)(`user.title`),subtitle:B(i)(`user.subtitle`)},{extra:Z(()=>[m(B(K),{type:`primary`,onClick:R},{default:Z(()=>[F(ve(B(i)(`user.newUser`)),1)]),_:1})]),_:1},8,[`title`,`subtitle`]),m(B(a),{style:{"margin-top":`16px`}},{default:Z(()=>[m(B(Xe),{value:y.value,"onUpdate:value":t[0]||=e=>y.value=e,type:`line`},{default:Z(()=>[m(B(qe),{name:`users`,tab:B(i)(`user.tabUsers`)},{default:Z(()=>[m(B(j),{columns:O.value,data:u.value,loading:s.value,pagination:{page:d.value.page,pageSize:d.value.pageSize,itemCount:d.value.itemCount,showSizePicker:!0,pageSizes:[10,20,50]},"row-key":e=>e.id,"onUpdate:page":P,"onUpdate:pageSize":te},null,8,[`columns`,`data`,`loading`,`pagination`,`row-key`])]),_:1},8,[`tab`]),m(B(qe),{name:`roles`,tab:B(i)(`user.tabRoles`)},{default:Z(()=>[m(B(j),{columns:A.value,data:M,bordered:!1,"row-key":e=>e.name},null,8,[`columns`,`row-key`])]),_:1},8,[`tab`])]),_:1},8,[`value`])]),_:1}),m(B(xe),{show:f.value,"onUpdate:show":t[8]||=e=>f.value=e,preset:`dialog`,title:v.value?B(i)(`user.editUser`):B(i)(`user.newUser`),"positive-text":B(i)(`common.save`),"negative-text":B(i)(`common.cancel`),loading:h.value,onPositiveClick:ne},{default:Z(()=>[m(B(Se),{ref_key:`formRef`,ref:p,model:S.value,rules:T,"label-placement":`left`,"label-width":`120`},{default:Z(()=>[m(B(H),{label:B(i)(`user.username`),path:`username`},{default:Z(()=>[m(B(E),{value:S.value.username,"onUpdate:value":t[1]||=e=>S.value.username=e,placeholder:B(i)(`user.username`),disabled:!!v.value},null,8,[`value`,`placeholder`,`disabled`])]),_:1},8,[`label`]),v.value?(L(),D(B(H),{key:1,label:B(i)(`user.password`),path:`password`},{default:Z(()=>[m(B(E),{value:S.value.password,"onUpdate:value":t[3]||=e=>S.value.password=e,type:`password`,"show-password-on":`click`,placeholder:B(i)(`user.keepPassword`)},null,8,[`value`,`placeholder`])]),_:1},8,[`label`])):(L(),D(B(H),{key:0,label:B(i)(`user.password`),path:`password`},{default:Z(()=>[m(B(E),{value:S.value.password,"onUpdate:value":t[2]||=e=>S.value.password=e,type:`password`,"show-password-on":`click`,placeholder:B(i)(`user.password`)},null,8,[`value`,`placeholder`])]),_:1},8,[`label`])),m(B(H),{label:B(i)(`user.email`),path:`email`},{default:Z(()=>[m(B(E),{value:S.value.email,"onUpdate:value":t[4]||=e=>S.value.email=e,placeholder:B(i)(`user.email`)},null,8,[`value`,`placeholder`])]),_:1},8,[`label`]),m(B(H),{label:B(i)(`user.role`),path:`roles`},{default:Z(()=>[m(B(l),{value:S.value.roles,"onUpdate:value":t[5]||=e=>S.value.roles=e,options:x,multiple:``,placeholder:B(i)(`user.role`)},null,8,[`value`,`placeholder`])]),_:1},8,[`label`]),m(B(H),{label:B(i)(`user.isSuperuser`),path:`is_superuser`},{default:Z(()=>[m(B(n),{value:S.value.is_superuser,"onUpdate:value":t[6]||=e=>S.value.is_superuser=e},null,8,[`value`])]),_:1},8,[`label`]),m(B(H),{label:B(i)(`user.isActive`),path:`is_active`},{default:Z(()=>[m(B(n),{value:S.value.is_active,"onUpdate:value":t[7]||=e=>S.value.is_active=e},null,8,[`value`])]),_:1},8,[`label`])]),_:1},8,[`model`])]),_:1},8,[`show`,`title`,`positive-text`,`negative-text`,`loading`])]))}});export{tt as default};