import os
from sys import platform as _platform
from ctypes import *
import secrethub_pb2
import base64


def lib_extension():
    if _platform == 'darwin':
        return 'dylib'
    elif _platform == 'win32':
        return 'dll'
    else:
        return 'so'


lib = cdll.LoadLibrary(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'resources', _platform, 'libsecrethub.'+lib_extension()))

# define class GoString to map:
# C type struct { const char *p; GoInt n; }
class GoString(Structure):
    _fields_ = [("p", c_char_p), ("n", c_longlong)]

class Client():    
    def read(self, path):
        req = secrethub_pb2.ReadRequest()
        req.path = path

        reqStr = req.SerializeToString()
        encoded = base64.b64encode(reqStr)
        request = GoString(encoded, len(encoded))

        lib.Read.argtypes = [GoString]
        lib.Read.restype = c_char_p
        result = lib.Read(request)
        
        
        response = secrethub_pb2.ReadResponse()
        response.ParseFromString(base64.b64decode(result))
        return response