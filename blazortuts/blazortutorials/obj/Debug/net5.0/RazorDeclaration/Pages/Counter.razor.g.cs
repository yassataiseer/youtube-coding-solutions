// <auto-generated/>
#pragma warning disable 1591
#pragma warning disable 0414
#pragma warning disable 0649
#pragma warning disable 0169

namespace blazortutorials.Pages
{
    #line hidden
    using System;
    using System.Collections.Generic;
    using System.Linq;
    using System.Threading.Tasks;
    using Microsoft.AspNetCore.Components;
#nullable restore
#line 1 "/Users/yassa/vidscode/blazortuts/blazortutorials/_Imports.razor"
using System.Net.Http;

#line default
#line hidden
#nullable disable
#nullable restore
#line 2 "/Users/yassa/vidscode/blazortuts/blazortutorials/_Imports.razor"
using System.Net.Http.Json;

#line default
#line hidden
#nullable disable
#nullable restore
#line 3 "/Users/yassa/vidscode/blazortuts/blazortutorials/_Imports.razor"
using Microsoft.AspNetCore.Components.Forms;

#line default
#line hidden
#nullable disable
#nullable restore
#line 4 "/Users/yassa/vidscode/blazortuts/blazortutorials/_Imports.razor"
using Microsoft.AspNetCore.Components.Routing;

#line default
#line hidden
#nullable disable
#nullable restore
#line 5 "/Users/yassa/vidscode/blazortuts/blazortutorials/_Imports.razor"
using Microsoft.AspNetCore.Components.Web;

#line default
#line hidden
#nullable disable
#nullable restore
#line 6 "/Users/yassa/vidscode/blazortuts/blazortutorials/_Imports.razor"
using Microsoft.AspNetCore.Components.Web.Virtualization;

#line default
#line hidden
#nullable disable
#nullable restore
#line 7 "/Users/yassa/vidscode/blazortuts/blazortutorials/_Imports.razor"
using Microsoft.AspNetCore.Components.WebAssembly.Http;

#line default
#line hidden
#nullable disable
#nullable restore
#line 8 "/Users/yassa/vidscode/blazortuts/blazortutorials/_Imports.razor"
using Microsoft.JSInterop;

#line default
#line hidden
#nullable disable
#nullable restore
#line 9 "/Users/yassa/vidscode/blazortuts/blazortutorials/_Imports.razor"
using blazortutorials;

#line default
#line hidden
#nullable disable
#nullable restore
#line 10 "/Users/yassa/vidscode/blazortuts/blazortutorials/_Imports.razor"
using blazortutorials.Shared;

#line default
#line hidden
#nullable disable
    [Microsoft.AspNetCore.Components.RouteAttribute("/counter")]
    public partial class Counter : Microsoft.AspNetCore.Components.ComponentBase
    {
        #pragma warning disable 1998
        protected override void BuildRenderTree(Microsoft.AspNetCore.Components.Rendering.RenderTreeBuilder __builder)
        {
        }
        #pragma warning restore 1998
#nullable restore
#line 31 "/Users/yassa/vidscode/blazortuts/blazortutorials/Pages/Counter.razor"

    private int currentCount = 0;
    public LoginModel LoginModel = new();
    private dynamic validate;
    private void IncrementCount()
    {
        Console.WriteLine(currentCount);
        currentCount++;
    }
    private async void HandleValidSumbit(){
        //Console.WriteLine(LoginModel.Username);
        //Console.WriteLine(LoginModel.Password);
        await JSRuntime.InvokeVoidAsync("BlazorSetLocalStorage","Username:",LoginModel.Username);
        validate = await JSRuntime.InvokeAsync<string>("BlazorGetLocalStorage","Username:");
        Console.WriteLine(validate);
    }

#line default
#line hidden
#nullable disable
        [global::Microsoft.AspNetCore.Components.InjectAttribute] private IJSRuntime JSRuntime { get; set; }
    }
}
#pragma warning restore 1591