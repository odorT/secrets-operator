import React from "react";
import FindingItem from "./FindingItem";
import propTypes from "prop-types";

function FindingList({ findings, repoURL }) {
  return (
    <div className="rounded-lg shadow-lg card bg-base-100">
      <div className="card-body">
        <h2 className="text-3xl my-4 font-bold card-title">
          Latest Findings
        </h2>
        {findings.map((finding) => (
          <FindingItem key={finding.id} finding={finding} repoURL={repoURL} />
        ))}
      </div>
    </div>
  );
}

FindingList.propTypes = {
    findings: propTypes.array.isRequired,
}

export default FindingList
