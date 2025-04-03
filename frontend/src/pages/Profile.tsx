import { useAuth } from '../contexts/AuthContext';
import { User } from 'lucide-react';

export default function Profile() {
  const { user, logout } = useAuth();

  return (
    <div className="min-h-screen bg-gray-100 py-12 px-4">
      <div className="max-w-2xl mx-auto bg-white rounded-lg shadow-md">
        <div className="p-6">
          <div className="flex items-center justify-center mb-6">
            <User className="w-20 h-20 text-blue-600" />
          </div>
          <h1 className="text-3xl font-bold text-center mb-8">Profil Utilisateur</h1>
          <div className="space-y-4">
            <div className="border-b pb-4">
              <p className="text-gray-600">Nom d'utilisateur</p>
              <p className="text-xl font-semibold">{user?.username}</p>
            </div>
            <div className="border-b pb-4">
              <p className="text-gray-600">Email</p>
              <p className="text-xl font-semibold">{user?.email}</p>
            </div>
            <div className="border-b pb-4">
              <p className="text-gray-600">ID Utilisateur</p>
              <p className="text-xl font-mono">{user?.id}</p>
            </div>
          </div>
          <button
            onClick={logout}
            className="mt-8 w-full bg-red-600 text-white font-bold py-2 px-4 rounded hover:bg-red-700 transition duration-200"
          >
            Se déconnecter
          </button>
        </div>
      </div>
    </div>
  );
}