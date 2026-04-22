#        -----------+ Non-Concurrent Test +-----------
The result from the database is:  id1  
The result from the database is:  id2  
The result from the database is:  id3  
The result from the database is:  id4  
The result from the database is:  id5  
The result from the database is:  id6  
The result from the database is:  id7  
The result from the database is:  id8  
The result from the database is:  id9  
The result from the database is:  id10  

_Total Execution Time: 20.0287424s_  
The results are [id1 id2 id3 id4 id5 id6 id7 id8 id9 id10]
#        -------------+ Concurrent Test +-------------
The result from the database is:  id9  
The result from the database is:  id3  
The result from the database is:  id1  
The result from the database is:  id8  
The result from the database is:  id7  
The result from the database is:  id6  
The result from the database is:  id4  
The result from the database is:  id5  
The result from the database is:  id2  
The result from the database is:  id10  

_Total Execution Time: 2.0020174s_  
The results are [id9 id3 id1 id8 id7 id6 id4 id5 id2 id10]