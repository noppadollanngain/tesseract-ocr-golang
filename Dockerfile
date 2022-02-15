FROM golang:1.16


ENV TERM=xterm
ENV TZ=Asia/Bangkok

RUN mkdir -p /usr/app/src
WORKDIR /usr/app/src

RUN apt-get update -qq

# You need librariy files and headers of tesseract and leptonica.
# When you miss these or LD_LIBRARY_PATH is not set to them,
# you would face an error: "tesseract/baseapi.h: No such file or directory"
RUN apt-get install -y -qq libtesseract-dev libleptonica-dev

# In case you face TESSDATA_PREFIX error, you minght need to set env vars
# to specify the directory where "tessdata" is located.
ENV TESSDATA_PREFIX=/usr/share/tesseract-ocr/4.00/tessdata/

# Load languages.
# These {lang}.traineddata would b located under ${TESSDATA_PREFIX}/tessdata.
RUN apt-get install -y -qq \
  tesseract-ocr-eng \
  tesseract-ocr-deu \
  tesseract-ocr-jpn

RUN go get -u github.com/cosmtrek/air

COPY ./src/.air.toml ./src/go.mod ./src/go.sum ./
RUN go mod tidy
RUN go mod download

COPY ./src .