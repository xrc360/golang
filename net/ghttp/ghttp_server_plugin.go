package ghttp

// Plugin is the interface for server plugin.
type Plugin interface {
	Name() string            // Name returns the name of the plugin.
	Author() string          // Author returns the author of the plugin.
	Version() string         // Version returns the version of the plugin, like "v1.0.0".
	Description() string     // Description returns the description of the plugin.
	Install(s *Server) error // Install installs the plugin BEFORE the server starts.
	Remove() error           // Remove removes the plugin when server shuts down.
}

// Plugin adds plugin to the server.
func (s *Server) Plugin(plugin ...Plugin) {
	s.plugins = append(s.plugins, plugin...)
}
