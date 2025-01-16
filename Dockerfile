FROM ghcr.io/carapace-sh/shell-elvish

RUN echo "deb [trusted=yes] https://apt.fury.io/rsteube/ /" \
       >  /etc/apt/sources.list.d/fury.list

RUN apt-get update && apt-get install -y asciinema carapace-bin gh git tmux

RUN echo "eval (carapace _carapace|slurp)" \
      >> /root/.config/elvish/rc.elv

RUN mkdir -p /root/.config/carapace \
 && echo "freckles: carapace" \
      >> /root/.config/carapace/bridges.yaml

RUN echo "[credential \"https://github.com\"]\n\
        helper = !gh auth git-credential\n"\
      > /root/.gitconfig

RUN echo "#!/bin/sh\nexec tmux new-session ';' resize-window -x 80 -y 12" \
       > /usr/local/bin/tmux-mini \
 && chmod +x /usr/local/bin/tmux-mini

ENV PATH=/freckles/cmd/freckles:$PATH
