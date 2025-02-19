package email

var REGISTER_EMAIL_SUBJECT = "Verify Your Email - Complete Your Registration"
var REGISTER_EMAIL_BODY_TEMPLATE = `Dear Friend,

Thank you for registering with Cooking Master. Please verify your email address by clicking the link below:  

http://127.0.0.1:8888/usercenter/v1/user/varifyregister?token=%s

If you did not create this account, please ignore this email.  

Best regards,  
Cooking Master
`

var RESET_EMAIL_SUBJECT = "Password Reset Request"
var RESET_EMAIL_BODY_TEMPLATE = `Dear Friend, 

We received a request to reset your password. If you made this request, please click the link below to reset your password:  

http://127.0.0.1:8888/usercenter/v1/user/passwordreset?token=%s

If you did not request a password reset, please ignore this email.  

Thank you,  
Cooking Master
`
