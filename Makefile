TAILSCALE_IP=$(shell ip addr show dev tailscale0 | grep "inet " | awk '{print $$2}' | cut -d/ -f1)

dev-server-run-tailscale:
	hugo server --bind="$(TAILSCALE_IP)" --port=1313 --baseURL="http://$(TAILSCALE_IP)"
