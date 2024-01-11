package email

import (
	"log"
	"net/smtp"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEmail() *schema.Resource {
	return &schema.Resource{
		Create: resourceEmailCreate,
		Read:   resourceEmailRead,
		Update: resourceEmailUpdate,
		Delete: resourceEmailDelete,
		Schema: emailResourceSchema(),
	}

}

func emailResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"to": {
			Type:     schema.TypeString,
			Required: true,
		},
		"from": {
			Type:     schema.TypeString,
			Required: true,
		},
		"reply_to": { // Add this field
			Type:     schema.TypeString,
			Optional: true,
			Default:  "",
		},
		"subject": {
			Type:     schema.TypeString,
			Required: true,
		},
		"preamble": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "",
		},
		"body": {
			Type:     schema.TypeString,
			Required: true,
		},
		"smtp_server": {
			Type:     schema.TypeString,
			Required: true,
		},
		"smtp_port": {
			Type:     schema.TypeString,
			Required: true,
		},
		"smtp_username": {
			Type:     schema.TypeString,
			Required: true,
		},
		"smtp_password": {
			Type:      schema.TypeString,
			Required:  true,
			Sensitive: true,
		},
	}
}

func resourceEmailCreate(d *schema.ResourceData, m interface{}) error {
	to := d.Get("to").(string)
	from := d.Get("from").(string)
	replyTo := d.Get("reply_to").(string)
	subject := d.Get("subject").(string)
	preamble := d.Get("preamble").(string)
	body := d.Get("body").(string)
	smtpServer := d.Get("smtp_server").(string)
	smtpPort := d.Get("smtp_port").(string)
	smtpUsername := d.Get("smtp_username").(string)
	smtpPassword := d.Get("smtp_password").(string)

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Reply-To: " + replyTo + "\n" +
		"Subject: " + subject + "\n" +
		preamble + "\n\n" +
		body

	err := smtp.SendMail(smtpServer+":"+smtpPort,
		smtp.PlainAuth("", smtpUsername, smtpPassword, smtpServer),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}

	// Create unique ID using current timestamp
	timestamp := time.Now().Unix()
	d.SetId(to + " | " + subject + " | " + strconv.FormatInt(timestamp, 10))
	return resourceEmailRead(d, m)
}

func resourceEmailRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceEmailDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}

func resourceEmailUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceEmailRead(d, m)
}
