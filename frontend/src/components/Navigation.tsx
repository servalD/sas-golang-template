import { Link } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';
import { Home, User, Image, UserPlus } from 'lucide-react';

export default function Navigation() {
  const { isAuthenticated } = useAuth();

  return (
    <nav className="bg-gray-800 text-white">
      <div className="max-w-7xl mx-auto px-4">
        <div className="flex items-center justify-between h-16">
          <div className="flex items-center space-x-8">
            <Link to="/" className="flex items-center space-x-2 hover:text-blue-400">
              <Home className="w-5 h-5" />
              <span>Accueil</span>
            </Link>
            {isAuthenticated && (
              <Link to="/nft-collection" className="flex items-center space-x-2 hover:text-blue-400">
                <Image className="w-5 h-5" />
                <span>Collection NFT</span>
              </Link>
            )}
          </div>
          <div className="flex items-center space-x-4">
            {isAuthenticated ? (
              <Link to="/profile" className="flex items-center space-x-2 hover:text-blue-400">
                <User className="w-5 h-5" />
                <span>Profil</span>
              </Link>
            ) : (
              <>
                <Link to="/login" className="flex items-center space-x-2 hover:text-blue-400">
                  <User className="w-5 h-5" />
                  <span>Connexion</span>
                </Link>
                <Link to="/signup" className="flex items-center space-x-2 hover:text-blue-400">
                  <UserPlus className="w-5 h-5" />
                  <span>Inscription</span>
                </Link>
              </>
            )}
          </div>
        </div>
      </div>
    </nav>
  );
}