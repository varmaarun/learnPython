def normal(*args):
    for stuff in args:
        print (stuff)



def kwargs_eg(**kwargs):
    for key, value in kwargs.items():
        print(key)
        print(value)

my_list=[1,2,3,4,5,6,7]
normal(1,2,3,4)
normal(*my_list)


d_eg={'key1':'value1','key2':'value2','key3':'value3','key4':'value4'}
kwargs_eg(**d_eg)