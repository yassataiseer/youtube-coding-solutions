using System.ComponentModel.DataAnnotations;

public class LoginModel{
    [Required]
    public string Username{get;set;}
    [Required]
    public string Password{get;set;}
}

