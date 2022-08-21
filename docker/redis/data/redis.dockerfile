FROM redis:6.2.6

ADD insert.sh  /usr/local/insert.sh 
ADD initial_data.txt /usr/local/initial_data.txt

RUN sed -i -e 's/\r$//' /usr/local/insert.sh
RUN chmod +x /usr/local/insert.sh

CMD ["redis-cli"]