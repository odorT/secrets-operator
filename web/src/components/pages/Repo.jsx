import { FaCodepen, FaStore, FaUserFriends, FaUsers } from "react-icons/fa";
import React, { useContext, useEffect } from "react";
import { useParams, Link } from "react-router-dom";
import RepoContext from "../../context/repos/RepoContext";
import Spinner from "../layout/Spinner";
import FindingList from "../layout/secrets/FindingList";
import { getRepoAndFindings } from "../../context/repos/RepoActions";
import { ACTIONS } from "../../context/repos/RepoReducer";

function Repo() {
  const { repo, loading, findings, dispatch } = useContext(RepoContext);
  const params = useParams();

  useEffect(() => {
    dispatch({ type: ACTIONS.SET_LOADING });

    const getRepoData = async () => {
      const repoData = await getRepoAndFindings(params.repo);
      dispatch({ type: ACTIONS.GET_REPOS_AND_FINDINGS, payload: repoData });
    };

    getRepoData();
  }, [dispatch, params.repo]);

  const { repoName, repoId, repoURL } = repo;

  if (loading) {
    return <Spinner />;
  }
  return (
    <>
      <div className="w-full mx-auto lg:w-10/12">
        <div className="mb-4">
          <Link to="/" className="btn btn-ghost">
            Back To Search
          </Link>
        </div>

        <div className="grid grid-cols-1 xl:grid-cols-3 lg:grid-cols-3 md:grid-cols-3 mb-8 md:gap-8">
          <div className="custom-card-image mb-6 md:mb-0">
            <div className="rounded-lg shadow-xl card image-full">
              <figure>
                <img src={"https://about.gitlab.com/images/press/logo/png/gitlab-logo-500.png"} alt="gitlab" />
              </figure>
              <div className="card-body justify-end">
                <h2 className="card-title mb-0">{repoName}</h2>
                <p className="flex-grow-0">{repoId}</p>
              </div>
            </div>
          </div>

          <div className="col-span-2">
            <div className="mb-6">
              <h1 className="text-3xl card-title">
                {repoName}
                {/* <div className="ml-2 mr-1 badge badge-success">{repoName}</div> */}
                {/* {repoName && (
                  <div className="mx-1 badge badge-info">repoName</div>
                )} */}
              </h1>
              {/* <p>{repoName}</p> */}
              <div className="mt-4 card-actions">
                <a
                  href={repoURL}
                  target="_blank"
                  rel="noreferrer"
                  className="btn btn-outline"
                >
                  Visit Repository
                </a>
              </div>
            </div>

            {/* <div className="w-full rounded-lg shadow-md bg-base-100 stats">
              {repoName && (
                <div className="stat">
                  <div className="stat-title text-md">Location</div>
                  <div className="text-lg stat-value">{repoName}</div>
                </div>
              )}
              {repoName && (
                <div className="stat">
                  <div className="stat-title text-md">Website</div>
                  <div className="text-lg stat-value">
                    <a href={repoName} target="_blank" rel="noreferrer">
                      {repoName}
                    </a>
                  </div>
                </div>
              )}
              {repoName && (
                <div className="stat">
                  <div className="stat-title text-md">Twitter</div>
                  <div className="text-lg stat-value">
                    <a
                      href={`https://twitter.com/${repoName}`}
                      target="_blank"
                      rel="noreferrer"
                    >
                      {repoName}
                    </a>
                  </div>
                </div>
              )}
            </div> */}
          </div>
        </div>

        {/* <div className="w-full py-5 mb-6 rounded-lg shadow-md bg-base-100 stats">
          <div className="grid grid-cols-1 md:grid-cols-3">
            <div className="stat">
              <div className="stat-figure text-secondary">
                <FaUsers className="text-3xl md:text-5xl" />
              </div>
              <div className="stat-title pr-5">Followers</div>
              <div className="stat-value pr-5 text-3xl md:text-4xl">
                {repoName}
              </div>
            </div>

            <div className="stat">
              <div className="stat-figure text-secondary">
                <FaUserFriends className="text-3xl md:text-5xl" />
              </div>
              <div className="stat-title pr-5">Following</div>
              <div className="stat-value pr-5 text-3xl md:text-4xl">
                {repoName}
              </div>
            </div>

            <div className="stat">
              <div className="stat-figure text-secondary">
                <FaCodepen className="text-3xl md:text-5xl" />
              </div>
              <div className="stat-title pr-5">Public Repos</div>
              <div className="stat-value pr-5 text-3xl md:text-4xl">
                {repoName}
              </div>
            </div>

            <div className="stat">
              <div className="stat-figure text-secondary">
                <FaStore className="text-3xl md:text-5xl" />
              </div>
              <div className="stat-title pr-5">Public Gists</div>
              <div className="stat-value pr-5 text-3xl md:text-4xl">
                {repoName}
              </div>
            </div>
          </div>
        </div> */}

        <FindingList findings={findings} repoURL={repoURL} />
      </div>
    </>
  );
}

export default Repo