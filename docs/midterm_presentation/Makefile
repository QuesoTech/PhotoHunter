TARGET=pres

all: spell pdf clean

pdf:
	pdflatex $(TARGET)
#	bibtex $(TARGET)
#	pdflatex $(TARGET)
#	pdflatex $(TARGET)

spell:
	aspell -ttex -c $(TARGET).tex

clean:
	rm -f *.blg *.aux *.bbl *.brf *.log *.nav *.snm *.toc *.out
