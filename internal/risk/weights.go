package risk

var RiskWeights = map[string]int{

	// Domain
	"lookalike domain": 70,
	"Newly registered domain": 80,
	"Suspicious sender reputation": 60,
	"Suspicious URL redirection": 70,
	"Repeated malicious sender": 100,
	"URL found in reputation database": 120,

	// Campaign
	"Potential phishing campaign": 150,

	// QR
	"QR code phishing detected": 200,
	"QR URL extracted": 120,

	// Malware
	"Known malware hash detected": 150,

	// ZIP
	"ZIP attachment detected": 60,
	"RAR attachment detected": 60,
	"ZIP contains suspicious file": 100,
	"Password-protected ZIP suspected": 150,
	"ZIP contains executable file": 150,
	"ZIP contains PowerShell file": 150,
	"ZIP contains macro-enabled document": 120,
	"ZIP contains macro-enabled spreadsheet": 120,
	"Nested ZIP detected": 200,

	// Office
	"Macro-enabled Office document detected": 120,
	"Macro-enabled Excel document detected": 120,
	"Suspicious macro detected": 120,

	// VirusTotal
	"VirusTotal malicious hash detected": 300,
	"VirusTotal suspicious hash detected": 150,
	"VirusTotal malicious URL detected": 300,
	"VirusTotal suspicious URL detected": 150,

	// Header
	"Reply-To mismatch detected": 100,
	"Return-Path mismatch detected": 80,
	"Display name spoofing detected": 120,

	// BEC
	"BEC indicator detected": 200,

	// WHOIS / DNS
	"WHOIS suspicious TLD detected": 60,
	"DNS reputation hit": 150,

	// PDF
	"PDF phishing keyword detected": 80,
	"PDF attachment detected": 40,

	// YARA
	"YARA rule matched": 100,

	// Sandbox
	"Sandbox behavior": 150,
	"Sandbox IOC URL": 120,
	"Sandbox IOC Domain": 80,
	"Process Tree": 120,
	"Network Activity": 100,
	"Dropped File": 150,
	"Persistence Detected": 200,
	"Behavior Rule": 250,
	"Docker Analysis": 120,
	"Docker YARA": 200,
	"encoded powershell": 200,

	// URL
	"Shortened URL detected": 70,
	"Homograph domain detected": 150,

	// Thread
	"Potential thread hijacking detected": 120,
	"Thread hijack indicator": 80,

	// Threat Intel
	"Threat feed hit": 200,
	"Detonation:": 150,
	"PhishTank hit:": 200,

	// Generic
	"Suspicious TLD": 25,
	"Brand impersonation": 40,
	"IP-based URL": 35,
	"Suspicious attachment": 45,
	"Suspicious country": 30,
	"Known malicious domain": 60,
	"Known malicious IP": 70,
	"Impossible travel": 80,

	// Email Authentication
	"SPF FAIL": 10,
	"DKIM FAIL": 10,
	"DMARC FAIL": 10,
}