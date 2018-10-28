FROM ruby:2.5

RUN gem install watir

RUN apt-get update -y
RUN apt-get install -y curl wget unzip

RUN wget -q -O chrome64_66.0.3359.181.deb http://www.slimjetbrowser.com/chrome/lnx/chrome64_66.0.3359.181.deb
RUN apt install -y ./chrome64_66.0.3359.181.deb
RUN rm -f chrome64_66.0.3359.181.deb

ENV CD_DIR /chromedriver
RUN mkdir $CD_DIR

RUN CD_VER=2.40 && \
    wget -P $CD_DIR https://chromedriver.storage.googleapis.com/$CD_VER/chromedriver_linux64.zip

RUN unzip $CD_DIR/chromedriver_linux64.zip -d $CD_DIR
RUN mv $CD_DIR/chromedriver /bin

VOLUME [ "/yogoa" ]
WORKDIR /yogoa

ENTRYPOINT ["ruby"]
