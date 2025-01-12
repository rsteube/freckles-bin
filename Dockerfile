FROM ghcr.io/carapace-sh/shell-elvish

RUN echo "deb [trusted=yes] https://apt.fury.io/rsteube/ /" \
       >  /etc/apt/sources.list.d/fury.list

RUN apt-get update && apt-get install -y asciinema carapace-bin gh git 

RUN echo "eval (carapace _carapace|slurp)" \
      >> /root/.config/elvish/rc.elv

RUN mkdir -p /root/.config/carapace \
 && echo "freckles: carapace" \
      >> /root/.config/carapace/bridges.yaml

RUN echo "[credential \"https://github.com\"]\n\
        helper = !gh auth git-credential\n"\
      > /root/.gitconfig

ENV PATH=/freckles/cmd/freckles:$PATH
