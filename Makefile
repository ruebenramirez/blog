TAILSCALE_IP=$(shell ip addr show dev tailscale0 | grep "inet " | awk '{print $$2}' | cut -d/ -f1)
WG_IP=$(shell ip addr show dev wg0 | grep "inet " | awk '{print $$2}' | cut -d/ -f1)

dev-server-run-tailscale:
	hugo server --bind="$(TAILSCALE_IP)" --port=1313 --baseURL="http://$(TAILSCALE_IP)"

dev-server-run-wg:
	hugo server --bind="$(WG_IP)" --port=1313 --baseURL="http://$(WG_IP)"

local-server-run:
	hugo server -D --port=1313
