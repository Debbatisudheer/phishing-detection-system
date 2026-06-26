import "./ResultsPanel.css";

function ResultsPanel({

    results,

    loading,

}) {

    if (loading) {

        return (

            <div className="playground-card">

                <h2>Analysis Summary</h2>

                <div className="loading-box">

                    Analyzing...

                </div>

            </div>

        );

    }

    if (!results) {

        return (

            <div className="playground-card">

                <h2>Analysis Summary</h2>

                <div className="empty-box">

                    Select a sample or upload a file.

                </div>

            </div>

        );

    }

    const sandboxDone =
        results.sandbox_completed;

    const display =
        sandboxDone
            ? results.sandbox
            : results;

    const score =
        display.risk_score || 0;

    let badgeClass =
        "risk-low";

    if (score >= 800) {

        badgeClass =
            "risk-critical";

    } else if (score >= 500) {

        badgeClass =
            "risk-high";

    } else if (score >= 100) {

        badgeClass =
            "risk-medium";

    }

    return (

        <div className="playground-card">

            <h2>

                Analysis Summary

            </h2>

            <div className="risk-card">

                <h3>

                    {

                        sandboxDone

                        ?

                        "Overall Risk"

                        :

                        "Quick Analysis"

                    }

                </h3>

                <div className="progress">

                    <div

                        className="progress-fill"

                        style={{

                            width: `${Math.min(score,1000)/10}%`,

                        }}

                    />

                </div>

                <h1>

                    {score}

                </h1>

                <span className={badgeClass}>

                    {display.risk_level}

                </span>

            </div>

            <div className="section">

                <h3>

                    Sandbox Status

                </h3>

                <div className="item">

                    {

                        sandboxDone

                        ?

                        "✅ Sandbox Analysis Completed"

                        :

                        "🟡 Waiting for Sandbox Analysis..."

                    }

                </div>

            </div>

            <div className="section">

                <h3>

                    Final Verdict

                </h3>

                <div className="item">

                    {

                        sandboxDone

                        ?

                        display.verdict

                        :

                        "Pending Sandbox Result"

                    }

                </div>

            </div>

            <div className="section">

                <h3>

                    Detection Findings

                </h3>

                {

                    (display.findings || []).length === 0

                    ?

                    <div className="empty-item">

                        No Findings

                    </div>

                    :

                    (
    Array.isArray(display.findings)
        ? display.findings
        : (display.findings || "").split("\n")
).map(

                        (finding,index)=>(

                            <div

                                key={index}

                                className="item"

                            >

                                ✓ {finding}

                            </div>

                        )

                    )

                }

            </div>

            <div className="recommendation">

                <h3>

                    Recommendation

                </h3>

                {

                    display.verdict

                }

            </div>

        </div>

    );

}

export default ResultsPanel;