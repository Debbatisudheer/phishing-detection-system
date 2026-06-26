const sampleEmails = [

  {
    id: 1,
    type: "safe",
    title: "Gmail Welcome",
    from: "welcome@gmail.com",
    to: "employee@company.com",
    subject: "Welcome to Gmail",
    body:
      "Welcome to Gmail. Your account has been successfully created."
  },

  {
    id: 2,
    type: "safe",
    title: "HR Notification",
    from: "hr@company.com",
    to: "employee@company.com",
    subject: "Leave Approved",
    body:
      "Your leave request has been approved."
  },

  {
    id: 3,
    type: "safe",
    title: "Amazon Receipt",
    from: "orders@amazon.com",
    to: "user@gmail.com",
    subject: "Your Amazon Order",
    body:
      "Thank you for shopping with Amazon."
  },

  {
    id: 4,
    type: "phishing",
    title: "Microsoft 365 Login",
    from: "security@microsoft-login.com",
    to: "employee@company.com",
    subject: "Password Expiring",
    body:
      "Your Microsoft account will expire today. Login immediately. Visit https://microsoft-login.com"
  },

  {
    id: 5,
    type: "phishing",
    title: "CEO Fraud",
    from: "ceo.company@gmail.com",
    to: "finance@company.com",
    subject: "Urgent Wire Transfer",
    body:
      "Transfer $50,000 immediately. Keep this confidential."
  },

  {
    id: 6,
    type: "phishing",
    title: "Bank Phishing",
    from: "support@sbi-security.com",
    to: "customer@gmail.com",
    subject: "Verify Your Account",
    body:
      "Click here immediately: https://fake-sbi-login.com"
  },

  {
    id: 7,
    type: "phishing",
    title: "QR Phishing",
    from: "help@company.com",
    to: "employee@company.com",
    subject: "Scan QR Code",
    body:
      "Please scan the attached QR code to verify your account.",
    demoFile:
      "/demo/qr/qrrscan.png"
  },

  {
    id: 8,
    type: "malware",
    title: "ZIP Malware",
    from: "invoice@vendor.com",
    to: "employee@company.com",
    subject: "Invoice Attached",
    body:
      "Please open the attached ZIP invoice.",
    demoFile:
      "/demo/zip/finalzip.zip"
  },

  {
    id: 9,
    type: "malware",
    title: "PDF Malware",
    from: "billing@company.com",
    to: "employee@company.com",
    subject: "Updated Invoice",
    body:
      "Please review the attached PDF.",
    demoFile:
      "/demo/pdf/phishing.pdf"
  },

  {
    id: 10,
    type: "malware",
    title: "PowerShell Downloader",
    from: "security@company.com",
    to: "employee@company.com",
    subject: "Security Update",
    body:
      "Please run the attached PowerShell script.",
    demoFile:
      "/demo/powershell/advanced_encoded_test.ps1"
  },

  {
    id: 11,
    type: "malware",
    title: "DOCM",
    from: "security@company.com",
    to: "employee@company.com",
    subject: "Security Update",
    body:
      "Please open the attached file.",
    demoFile:
      "/demo/office/real_invoicee.docm"
  },

  {
    id: 12,
    type: "safe",
    title: "Custom Email",
    from: "",
    to: "",
    subject: "",
    body: ""
  }

];

export default sampleEmails;