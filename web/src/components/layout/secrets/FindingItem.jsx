import { FaEye, FaInfo, FaGit, FaLink, FaStar, FaUtensils } from "react-icons/fa";
import PropTypes from "prop-types";

function FindingItem({ finding, repoURL }) {
  const {
    Description,
    StartLine,
    EndLine,
    StartColumn,
    EndColumn,
    Match,
    Secret,
    File,
    Commit,
    Entropy,
    Author,
    Email,
    Date,
    Message,
    Tags,
    RuleID,
    Fingerprint,
  } = finding;


  let commitURL = repoURL + "/-/commit/" + Commit

  return (
    <div className="mb-2 rounded-md card bg-base-200 hover:bg-base-300">
      <div className="card-body">
        <h3 className="mb-2 text-xl font-semibold">
          <a href={commitURL}>
            <FaLink className="inline mr-1" /> {File}
          </a>
        </h3>

        <p className="mb-3">{Description}</p>
        
        <div>
          <h3>Secret: {Secret}</h3>
        </div>
        
        <div className="mr-2">
          {/* <div className="mr-2 badge badge-info badge-lg">
            <FaGit className="mr-2" />Commit: {Commit}
          </div>
          <div className="mr-2 badge badge-success badge-lg">
            <FaStar className="mr-2" /> {Author}
          </div>
          <div className="mr-2 badge badge-error badge-lg">
            <FaInfo className="mr-2" /> {Email}
          </div>
          <div className="mr-2 badge badge-warning badge-lg">
            <FaUtensils className="mr-2" /> {Date}
          </div> */}
          <p>Commit: {Commit}</p>
          <p>Entropy: {Entropy}</p>
          <p>Author: {Author}</p>
          <p>Email: {Email}</p>
          <p>Date: {Date}</p>
          <p>Rule ID: {RuleID}</p>
          <p>FingerPrint: {Fingerprint}</p>
        </div>
      </div>
    </div>
  );
}

FindingItem.propTypes = {
  finding: PropTypes.object.isRequired,
}

export default FindingItem
