import React, { useState, useEffect } from 'react';
import { Calendar, Clock, ArrowLeft, ExternalLink } from 'lucide-react';

interface Discussion {
  id: string;
  title: string;
  body: string;
  createdAt: string;
  updatedAt: string;
  author: {
    login: string;
    avatarUrl: string;
  };
  url: string;
  category: {
    name: string;
  };
}

interface BlogPageProps {
  onBack: () => void;
}

const BlogPage: React.FC<BlogPageProps> = ({ onBack }) => {
  const [discussions, setDiscussions] = useState<Discussion[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetchDiscussions();
  }, []);

  const fetchDiscussions = async () => {
    try {
      setLoading(true);
      setError(null);
      
      // GitHub GraphQL query to fetch discussions in announcements category
      const query = `
        query {
          repository(owner: "quaadgras", name: "graphics.gd") {
            discussions(first: 20, categorySlug: "announcements", orderBy: {field: CREATED_AT, direction: DESC}) {
              nodes {
                id
                title
                body
                createdAt
                updatedAt
                author {
                  login
                  avatarUrl
                }
                url
                category {
                  name
                }
              }
            }
          }
        }
      `;

      const response = await fetch('https://api.github.com/graphql', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ query }),
      });

      if (!response.ok) {
        throw new Error('Failed to fetch discussions');
      }

      const data = await response.json();
      
      if (data.errors) {
        throw new Error(data.errors[0].message);
      }

      setDiscussions(data.data.repository.discussions.nodes);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'An error occurred');
      console.error('Error fetching discussions:', err);
    } finally {
      setLoading(false);
    }
  };

  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'long',
      day: 'numeric',
    });
  };

  const truncateText = (text: string, maxLength: number = 200) => {
    if (text.length <= maxLength) return text;
    return text.substring(0, maxLength) + '...';
  };

  return (
    <div className="min-h-screen bg-white">
      {/* Header */}
      <div className="bg-gradient-to-br from-blue-50 via-white to-teal-50 border-b border-gray-200">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
          <button
            onClick={onBack}
            className="inline-flex items-center space-x-2 text-blue-600 hover:text-blue-700 mb-6 transition-colors"
          >
            <ArrowLeft className="h-4 w-4" />
            <span>Back to Home</span>
          </button>
          
          <div className="text-center">
            <h1 className="text-4xl sm:text-5xl font-bold text-gray-900 mb-4">
              Blog & Announcements
            </h1>
            <p className="text-xl text-gray-600 max-w-3xl mx-auto">
              Stay up to date with the latest news, updates, and announcements for graphics.gd
            </p>
          </div>
        </div>
      </div>

      {/* Content */}
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
        {loading && (
          <div className="text-center py-12">
            <div className="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
            <p className="mt-4 text-gray-600">Loading announcements...</p>
          </div>
        )}

        {error && (
          <div className="text-center py-12">
            <div className="bg-red-50 border border-red-200 rounded-lg p-6 max-w-md mx-auto">
              <h3 className="text-lg font-medium text-red-800 mb-2">Error</h3>
              <p className="text-red-600">{error}</p>
              <button
                onClick={fetchDiscussions}
                className="mt-4 bg-red-600 text-white px-4 py-2 rounded-lg hover:bg-red-700 transition-colors"
              >
                Try Again
              </button>
            </div>
          </div>
        )}

        {!loading && !error && discussions.length === 0 && (
          <div className="text-center py-12">
            <h3 className="text-lg font-medium text-gray-900 mb-2">No announcements yet</h3>
            <p className="text-gray-600">Check back later for updates!</p>
          </div>
        )}

        {!loading && !error && discussions.length > 0 && (
          <div className="space-y-8">
            {discussions.map((discussion) => (
              <article
                key={discussion.id}
                className="bg-white border border-gray-200 rounded-2xl p-8 hover:shadow-lg transition-shadow"
              >
                <div className="flex items-start justify-between mb-4">
                  <div className="flex items-center space-x-3">
                    <img
                      src={discussion.author.avatarUrl}
                      alt={discussion.author.login}
                      className="w-10 h-10 rounded-full"
                    />
                    <div>
                      <p className="font-medium text-gray-900">
                        {discussion.author.login}
                      </p>
                      <div className="flex items-center space-x-4 text-sm text-gray-500">
                        <div className="flex items-center space-x-1">
                          <Calendar className="h-4 w-4" />
                          <span>{formatDate(discussion.createdAt)}</span>
                        </div>
                        {discussion.updatedAt !== discussion.createdAt && (
                          <div className="flex items-center space-x-1">
                            <Clock className="h-4 w-4" />
                            <span>Updated {formatDate(discussion.updatedAt)}</span>
                          </div>
                        )}
                      </div>
                    </div>
                  </div>
                  <span className="inline-flex items-center px-3 py-1 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
                    {discussion.category.name}
                  </span>
                </div>

                <h2 className="text-2xl font-bold text-gray-900 mb-4">
                  {discussion.title}
                </h2>

                <div className="prose prose-gray max-w-none mb-6">
                  <p className="text-gray-600 leading-relaxed">
                    {truncateText(discussion.body)}
                  </p>
                </div>

                <div className="flex items-center justify-between">
                  <a
                    href={discussion.url}
                    target="_blank"
                    rel="noopener noreferrer"
                    className="inline-flex items-center space-x-2 text-blue-600 hover:text-blue-700 font-medium transition-colors"
                  >
                    <span>Read full discussion</span>
                    <ExternalLink className="h-4 w-4" />
                  </a>
                </div>
              </article>
            ))}
          </div>
        )}

        {!loading && !error && discussions.length > 0 && (
          <div className="text-center mt-12">
            <a
              href="https://github.com/quaadgras/graphics.gd/discussions/categories/announcements"
              target="_blank"
              rel="noopener noreferrer"
              className="inline-flex items-center space-x-2 bg-gray-900 text-white px-6 py-3 rounded-lg hover:bg-gray-800 transition-colors"
            >
              <span>View all announcements on GitHub</span>
              <ExternalLink className="h-4 w-4" />
            </a>
          </div>
        )}
      </div>
    </div>
  );
};

export default BlogPage;