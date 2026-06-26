import { useEffect, useState } from "react";

function EmailEditor({

  selectedEmail,

  onAnalyze,

}) {

  const [email, setEmail] =
    useState({

      from: "",

      to: "",

      subject: "",

      body: "",

    });

  useEffect(() => {

    if (selectedEmail) {

      setEmail(selectedEmail);

    }

  }, [selectedEmail]);

  return (

    <div className="playground-card">

      <div className="email-header">

        <div className="avatar">

          📧

        </div>

        <div>

          <h2>

            Email Preview

          </h2>

          <p>

            Enterprise Email Analyzer

          </p>

        </div>

      </div>

      <div className="email-info">

        <div>

          <strong>

            From

          </strong>

          <input

            value={email.from}

            onChange={(e)=>

              setEmail({

                ...email,

                from:e.target.value,

              })

            }

          />

        </div>

        <div>

          <strong>

            To

          </strong>

          <input

            value={email.to}

            onChange={(e)=>

              setEmail({

                ...email,

                to:e.target.value,

              })

            }

          />

        </div>

        <div>

          <strong>

            Subject

          </strong>

          <input

            value={email.subject}

            onChange={(e)=>

              setEmail({

                ...email,

                subject:e.target.value,

              })

            }

          />

        </div>

      </div>

      <div
        className="email-body"
      >

        <strong>

          Email Body

        </strong>

        <textarea

          rows={16}

          value={email.body}

          onChange={(e)=>

            setEmail({

              ...email,

              body:e.target.value,

            })

          }

        />

      </div>

      <div className="editor-buttons">

        <button

          onClick={()=>

            onAnalyze(email)

          }

        >

          Analyze Email

        </button>

        <button

          onClick={()=>

            setEmail({

              from:"",

              to:"",

              subject:"",

              body:"",

            })

          }

        >

          Clear

        </button>

      </div>

    </div>

  );

}

export default EmailEditor;