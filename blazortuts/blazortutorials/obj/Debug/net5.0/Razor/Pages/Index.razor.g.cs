#pragma checksum "/Users/yassa/vidscode/blazortuts/blazortutorials/Pages/Index.razor" "{ff1816ec-aa5e-4d10-87f7-6f4963833460}" "788ccf89169519d20855146a53fef8376499995f"
// <auto-generated/>
#pragma warning disable 1591
namespace blazortutorials.Pages
{
    #line hidden
    using System;
    using System.Collections.Generic;
    using System.Linq;
    using System.Threading.Tasks;
    using Microsoft.AspNetCore.Components;
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
#nullable restore
#line 35 "/Users/yassa/vidscode/blazortuts/blazortutorials/Pages/Index.razor"
using System.Net.Http;

#line default
#line hidden
#nullable disable
#nullable restore
#line 36 "/Users/yassa/vidscode/blazortuts/blazortutorials/Pages/Index.razor"
using System.Text;

#line default
#line hidden
#nullable disable
#nullable restore
#line 37 "/Users/yassa/vidscode/blazortuts/blazortutorials/Pages/Index.razor"
using System.Net.Http.Json;

#line default
#line hidden
#nullable disable
#nullable restore
#line 38 "/Users/yassa/vidscode/blazortuts/blazortutorials/Pages/Index.razor"
using System.Web;

#line default
#line hidden
#nullable disable
#nullable restore
#line 39 "/Users/yassa/vidscode/blazortuts/blazortutorials/Pages/Index.razor"
using Newtonsoft.Json.Linq;

#line default
#line hidden
#nullable disable
#nullable restore
#line 40 "/Users/yassa/vidscode/blazortuts/blazortutorials/Pages/Index.razor"
using Newtonsoft.Json;

#line default
#line hidden
#nullable disable
    [Microsoft.AspNetCore.Components.RouteAttribute("/")]
    public partial class Index : Microsoft.AspNetCore.Components.ComponentBase
    {
        #pragma warning disable 1998
        protected override void BuildRenderTree(Microsoft.AspNetCore.Components.Rendering.RenderTreeBuilder __builder)
        {
            __builder.AddMarkupContent(0, "<h1>Hello, world!</h1>\r\n\r\n");
            __builder.OpenComponent<blazortutorials.Shared.Tooltip>(1);
            __builder.AddAttribute(2, "Text", "Tooltip Implementation");
            __builder.AddAttribute(3, "ChildContent", (Microsoft.AspNetCore.Components.RenderFragment)((__builder2) => {
                __builder2.AddContent(4, "Welcome to your new app.");
            }
            ));
            __builder.CloseComponent();
            __builder.AddMarkupContent(5, "\r\n\r\n");
            __builder.OpenComponent<blazortutorials.Shared.SurveyPrompt>(6);
            __builder.AddAttribute(7, "Title", "How is Blazor working for you?");
            __builder.CloseComponent();
            __builder.AddMarkupContent(8, "\r\n\r\n");
            __builder.OpenElement(9, "table");
            __builder.AddAttribute(10, "class", "table");
            __builder.AddMarkupContent(11, "<thead><tr><th>Name</th>\r\n            <th>Barcode</th>\r\n            <th>Price</th>\r\n            <th>Quantity</th>\r\n            <th>Status</th></tr></thead>\r\n    ");
            __builder.OpenElement(12, "tbody");
#nullable restore
#line 20 "/Users/yassa/vidscode/blazortuts/blazortutorials/Pages/Index.razor"
         foreach(var Current_Inventory in InventoryData){

#line default
#line hidden
#nullable disable
            __builder.OpenElement(13, "tr");
            __builder.OpenElement(14, "td");
            __builder.AddContent(15, 
#nullable restore
#line 22 "/Users/yassa/vidscode/blazortuts/blazortutorials/Pages/Index.razor"
                     Current_Inventory.Item

#line default
#line hidden
#nullable disable
            );
            __builder.CloseElement();
            __builder.AddMarkupContent(16, "\r\n                ");
            __builder.OpenElement(17, "td");
            __builder.AddContent(18, 
#nullable restore
#line 23 "/Users/yassa/vidscode/blazortuts/blazortutorials/Pages/Index.razor"
                     Current_Inventory.Barcode

#line default
#line hidden
#nullable disable
            );
            __builder.CloseElement();
            __builder.AddMarkupContent(19, "\r\n                ");
            __builder.OpenElement(20, "td");
            __builder.AddContent(21, 
#nullable restore
#line 24 "/Users/yassa/vidscode/blazortuts/blazortutorials/Pages/Index.razor"
                     Current_Inventory.Price

#line default
#line hidden
#nullable disable
            );
            __builder.CloseElement();
            __builder.AddMarkupContent(22, "\r\n                ");
            __builder.OpenElement(23, "td");
            __builder.AddContent(24, 
#nullable restore
#line 25 "/Users/yassa/vidscode/blazortuts/blazortutorials/Pages/Index.razor"
                      Current_Inventory.Quantity

#line default
#line hidden
#nullable disable
            );
            __builder.CloseElement();
            __builder.AddMarkupContent(25, "\r\n                ");
            __builder.OpenElement(26, "td");
            __builder.AddContent(27, 
#nullable restore
#line 26 "/Users/yassa/vidscode/blazortuts/blazortutorials/Pages/Index.razor"
                      Current_Inventory.Status

#line default
#line hidden
#nullable disable
            );
            __builder.CloseElement();
            __builder.CloseElement();
#nullable restore
#line 28 "/Users/yassa/vidscode/blazortuts/blazortutorials/Pages/Index.razor"
        }

#line default
#line hidden
#nullable disable
            __builder.CloseElement();
            __builder.CloseElement();
        }
        #pragma warning restore 1998
#nullable restore
#line 40 "/Users/yassa/vidscode/blazortuts/blazortutorials/Pages/Index.razor"
                          
    public List<InventoryModel> InventoryData = new();

    protected override async Task OnInitializedAsync(){
        using var client = new HttpClient();
        var result = await client.GetStringAsync("http://0.0.0.0:800/Inventory/grab_Inventory");
        JArray data = JArray.Parse(result);
        foreach(dynamic obj in data){
            InventoryData.Add(new InventoryModel(){
                Barcode = obj.Barcode,
                Item  = obj.Item,
                Price = obj.Price,
                Quantity = obj.Quantity,
                Status = obj.Status
            });

        }
        StateHasChanged();
    }

#line default
#line hidden
#nullable disable
    }
}
#pragma warning restore 1591