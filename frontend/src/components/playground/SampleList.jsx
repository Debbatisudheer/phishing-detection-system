import { useState } from "react";
import sampleEmails from "../../data/sampleEmails";

function SampleList({

  onSelect,

}) {

  const [selected, setSelected] =
    useState(null);

  const [search, setSearch] =
    useState("");

  const chooseSample =
    (sample) => {

      setSelected(
        sample.id,
      );

      onSelect(
        sample,
      );

    };

  const filteredSamples =
    sampleEmails.filter(

      (sample) =>

        sample.title
          .toLowerCase()
          .includes(
            search.toLowerCase(),
          ),

    );

  const renderSection =
    (
      title,
      type,
      icon,
    ) => (

      <>

        <h3>

          {icon} {title}

        </h3>

        {

          filteredSamples

            .filter(

              sample =>
                sample.type === type,

            )

            .map(

              sample => (

                <button

                  key={
                    sample.id
                  }

                  className={

                    selected ===
                    sample.id

                      ?

                      "sample-button active"

                      :

                      "sample-button"

                  }

                  onClick={() =>

                    chooseSample(
                      sample,
                    )

                  }

                >

                  <div>

                    <strong>

                      {sample.title}

                    </strong>

                  </div>

                  <small>

                    {

                      sample.demoFile

                        ?

                        "Built-in Demo File"

                        :

                        "Email Sample"

                    }

                  </small>

                </button>

              ),

            )

        }

      </>

    );

  return (

    <div
      className="playground-card"
    >

      <h2>

        Demo Library

      </h2>

      <input

        className="sample-search"

        placeholder="Search Samples..."

        value={search}

        onChange={(e) =>

          setSearch(
            e.target.value,
          )

        }

      />

      {

        renderSection(

          "Legitimate",

          "safe",

          "✅",

        )

      }

      {

        renderSection(

          "Phishing",

          "phishing",

          "🎣",

        )

      }

      {

        renderSection(

          "Malware",

          "malware",

          "☣️",

        )

      }

    </div>

  );

}

export default SampleList;