import{An as e,Bn as t,Cr as n,Dr as r,Er as i,Ir as a,Sr as o,Tr as s,Ur as c,Vn as l,Wr as u,Xt as d,ar as f,br as p,dn as m,dr as h,jn as g,pr as _,wr as v,xr as y}from"./index-Dmrq63M5.js";function b(e,t=`default`,n=[]){let{children:r}=e;if(typeof r==`object`&&r&&!Array.isArray(r)){let e=r[t];if(typeof e==`function`)return e()}return n}var x={thPaddingBorderedSmall:`8px 12px`,thPaddingBorderedMedium:`12px 16px`,thPaddingBorderedLarge:`16px 24px`,thPaddingSmall:`0`,thPaddingMedium:`0`,thPaddingLarge:`0`,tdPaddingBorderedSmall:`8px 12px`,tdPaddingBorderedMedium:`12px 16px`,tdPaddingBorderedLarge:`16px 24px`,tdPaddingSmall:`0 0 8px 0`,tdPaddingMedium:`0 0 12px 0`,tdPaddingLarge:`0 0 16px 0`};function S(e){let{tableHeaderColor:t,textColor2:n,textColor1:r,cardColor:i,modalColor:a,popoverColor:o,dividerColor:s,borderRadius:c,fontWeightStrong:l,lineHeight:u,fontSizeSmall:d,fontSizeMedium:f,fontSizeLarge:p}=e;return Object.assign(Object.assign({},x),{lineHeight:u,fontSizeSmall:d,fontSizeMedium:f,fontSizeLarge:p,titleTextColor:r,thColor:_(i,t),thColorModal:_(a,t),thColorPopover:_(o,t),thTextColor:r,thFontWeight:l,tdTextColor:n,tdColor:i,tdColorModal:a,tdColorPopover:o,borderColor:_(i,s),borderColorModal:_(a,s),borderColorPopover:_(o,s),borderRadius:c})}var C={name:`Descriptions`,common:d,self:S},w=p([y(`descriptions`,{fontSize:`var(--n-font-size)`},[y(`descriptions-separator`,`
 display: inline-block;
 margin: 0 8px 0 2px;
 `),y(`descriptions-table-wrapper`,[y(`descriptions-table`,[y(`descriptions-table-row`,[y(`descriptions-table-header`,{padding:`var(--n-th-padding)`}),y(`descriptions-table-content`,{padding:`var(--n-td-padding)`})])])]),v(`bordered`,[y(`descriptions-table-wrapper`,[y(`descriptions-table`,[y(`descriptions-table-row`,[p(`&:last-child`,[y(`descriptions-table-content`,{paddingBottom:0})])])])])]),n(`left-label-placement`,[y(`descriptions-table-content`,[p(`> *`,{verticalAlign:`top`})])]),n(`left-label-align`,[p(`th`,{textAlign:`left`})]),n(`center-label-align`,[p(`th`,{textAlign:`center`})]),n(`right-label-align`,[p(`th`,{textAlign:`right`})]),n(`bordered`,[y(`descriptions-table-wrapper`,`
 border-radius: var(--n-border-radius);
 overflow: hidden;
 background: var(--n-merged-td-color);
 border: 1px solid var(--n-merged-border-color);
 `,[y(`descriptions-table`,[y(`descriptions-table-row`,[p(`&:not(:last-child)`,[y(`descriptions-table-content`,{borderBottom:`1px solid var(--n-merged-border-color)`}),y(`descriptions-table-header`,{borderBottom:`1px solid var(--n-merged-border-color)`})]),y(`descriptions-table-header`,`
 font-weight: 400;
 background-clip: padding-box;
 background-color: var(--n-merged-th-color);
 `,[p(`&:not(:last-child)`,{borderRight:`1px solid var(--n-merged-border-color)`})]),y(`descriptions-table-content`,[p(`&:not(:last-child)`,{borderRight:`1px solid var(--n-merged-border-color)`})])])])])]),y(`descriptions-header`,`
 font-weight: var(--n-th-font-weight);
 font-size: 18px;
 transition: color .3s var(--n-bezier);
 line-height: var(--n-line-height);
 margin-bottom: 16px;
 color: var(--n-title-text-color);
 `),y(`descriptions-table-wrapper`,`
 transition:
 background-color .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
 `,[y(`descriptions-table`,`
 width: 100%;
 border-collapse: separate;
 border-spacing: 0;
 box-sizing: border-box;
 `,[y(`descriptions-table-row`,`
 box-sizing: border-box;
 transition: border-color .3s var(--n-bezier);
 `,[y(`descriptions-table-header`,`
 font-weight: var(--n-th-font-weight);
 line-height: var(--n-line-height);
 display: table-cell;
 box-sizing: border-box;
 color: var(--n-th-text-color);
 transition:
 color .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
 `),y(`descriptions-table-content`,`
 vertical-align: top;
 line-height: var(--n-line-height);
 display: table-cell;
 box-sizing: border-box;
 color: var(--n-td-text-color);
 transition:
 color .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
 `,[o(`content`,`
 transition: color .3s var(--n-bezier);
 display: inline-block;
 color: var(--n-td-text-color);
 `)]),o(`label`,`
 font-weight: var(--n-th-font-weight);
 transition: color .3s var(--n-bezier);
 display: inline-block;
 margin-right: 14px;
 color: var(--n-th-text-color);
 `)])])])]),y(`descriptions-table-wrapper`,`
 --n-merged-th-color: var(--n-th-color);
 --n-merged-td-color: var(--n-td-color);
 --n-merged-border-color: var(--n-border-color);
 `),i(y(`descriptions-table-wrapper`,`
 --n-merged-th-color: var(--n-th-color-modal);
 --n-merged-td-color: var(--n-td-color-modal);
 --n-merged-border-color: var(--n-border-color-modal);
 `)),r(y(`descriptions-table-wrapper`,`
 --n-merged-th-color: var(--n-th-color-popover);
 --n-merged-td-color: var(--n-td-color-popover);
 --n-merged-border-color: var(--n-border-color-popover);
 `))]),T=`DESCRIPTION_ITEM_FLAG`;function E(e){return typeof e==`object`&&e&&!Array.isArray(e)?e.type&&e.type.DESCRIPTION_ITEM_FLAG:!1}var D=c({name:`Descriptions`,props:Object.assign(Object.assign({},m.props),{title:String,column:{type:Number,default:3},columns:Number,labelPlacement:{type:String,default:`top`},labelAlign:{type:String,default:`left`},separator:{type:String,default:`:`},size:String,bordered:Boolean,labelClass:String,labelStyle:[Object,String],contentClass:String,contentStyle:[Object,String]}),slots:Object,setup(t){let{mergedClsPrefixRef:n,inlineThemeDisabled:r,mergedComponentPropsRef:i}=g(t),o=a(()=>t.size||i?.value?.Descriptions?.size||`medium`),c=m(`Descriptions`,`-descriptions`,w,C,t,n),l=a(()=>{let{bordered:e}=t,n=o.value,{common:{cubicBezierEaseInOut:r},self:{titleTextColor:i,thColor:a,thColorModal:l,thColorPopover:u,thTextColor:d,thFontWeight:f,tdTextColor:p,tdColor:m,tdColorModal:h,tdColorPopover:g,borderColor:_,borderColorModal:v,borderColorPopover:y,borderRadius:b,lineHeight:x,[s(`fontSize`,n)]:S,[s(e?`thPaddingBordered`:`thPadding`,n)]:C,[s(e?`tdPaddingBordered`:`tdPadding`,n)]:w}}=c.value;return{"--n-title-text-color":i,"--n-th-padding":C,"--n-td-padding":w,"--n-font-size":S,"--n-bezier":r,"--n-th-font-weight":f,"--n-line-height":x,"--n-th-text-color":d,"--n-td-text-color":p,"--n-th-color":a,"--n-th-color-modal":l,"--n-th-color-popover":u,"--n-td-color":m,"--n-td-color-modal":h,"--n-td-color-popover":g,"--n-border-radius":b,"--n-border-color":_,"--n-border-color-modal":v,"--n-border-color-popover":y}}),u=r?e(`descriptions`,a(()=>{let e=``,{bordered:n}=t;return n&&(e+=`a`),e+=o.value[0],e}),l,t):void 0;return{mergedClsPrefix:n,cssVars:r?void 0:l,themeClass:u?.themeClass,onRender:u?.onRender,compitableColumn:f(t,[`columns`,`column`]),inlineThemeDisabled:r,mergedSize:o}},render(){let e=this.$slots.default,n=e?l(e()):[];n.length;let{contentClass:r,labelClass:i,compitableColumn:a,labelPlacement:o,labelAlign:s,mergedSize:c,bordered:d,title:f,cssVars:p,mergedClsPrefix:m,separator:g,onRender:_}=this;_?.();let v=n.filter(e=>E(e)),y=v.reduce((e,t,n)=>{let s=t.props||{},c=v.length-1===n,l=[`label`in s?s.label:b(t,`label`)],f=[b(t)],p=s.span||1,h=e.span;e.span+=p;let _=s.labelStyle||s[`label-style`]||this.labelStyle,y=s.contentStyle||s[`content-style`]||this.contentStyle;if(o===`left`)d?e.row.push(u(`th`,{class:[`${m}-descriptions-table-header`,i],colspan:1,style:_},l),u(`td`,{class:[`${m}-descriptions-table-content`,r],colspan:c?(a-h)*2+1:p*2-1,style:y},f)):e.row.push(u(`td`,{class:`${m}-descriptions-table-content`,colspan:c?(a-h)*2:p*2},u(`span`,{class:[`${m}-descriptions-table-content__label`,i],style:_},[...l,g&&u(`span`,{class:`${m}-descriptions-separator`},g)]),u(`span`,{class:[`${m}-descriptions-table-content__content`,r],style:y},f)));else{let t=c?(a-h)*2:p*2;e.row.push(u(`th`,{class:[`${m}-descriptions-table-header`,i],colspan:t,style:_},l)),e.secondRow.push(u(`td`,{class:[`${m}-descriptions-table-content`,r],colspan:t,style:y},f))}return(e.span>=a||c)&&(e.span=0,e.row.length&&(e.rows.push(e.row),e.row=[]),o!==`left`&&e.secondRow.length&&(e.rows.push(e.secondRow),e.secondRow=[])),e},{span:0,row:[],secondRow:[],rows:[]}).rows.map(e=>u(`tr`,{class:`${m}-descriptions-table-row`},e));return u(`div`,{style:p,class:[`${m}-descriptions`,this.themeClass,`${m}-descriptions--${o}-label-placement`,`${m}-descriptions--${s}-label-align`,`${m}-descriptions--${c}-size`,d&&`${m}-descriptions--bordered`]},f||this.$slots.header?u(`div`,{class:`${m}-descriptions-header`},f||t(this,`header`)):null,u(`div`,{class:`${m}-descriptions-table-wrapper`},u(`table`,{class:`${m}-descriptions-table`},u(`tbody`,null,o===`top`&&u(`tr`,{class:`${m}-descriptions-table-row`,style:{visibility:`collapse`}},h(a*2,u(`td`,null))),y))))}}),O={label:String,span:{type:Number,default:1},labelClass:String,labelStyle:[Object,String],contentClass:String,contentStyle:[Object,String]},k=c({name:`DescriptionsItem`,[T]:!0,props:O,slots:Object,render(){return null}});export{D as n,k as t};